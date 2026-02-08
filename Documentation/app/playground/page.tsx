"use client";

import { useState } from "react";
import { Play, RotateCcw, ChevronDown } from "lucide-react";
import { EXAMPLES } from "./examples";

export default function Playground() {
  const [selectedExample, setSelectedExample] = useState("hello.bang");
  const [code, setCode] = useState(EXAMPLES["hello.bang"].code);
  const [output, setOutput] = useState<string[]>([]);
  const [isRunning, setIsRunning] = useState(false);

  const handleExampleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const key = e.target.value;
    setSelectedExample(key);
    setCode(EXAMPLES[key as keyof typeof EXAMPLES].code);
    setOutput([]);
  };

  const handleRun = async () => {
    setIsRunning(true);
    setOutput([]);

    // Simulating compilation delay
    setTimeout(() => {
      // 1. Check for exact example match
      const matchedExample = Object.values(EXAMPLES).find(ex => ex.code.trim() === code.trim());

      if (matchedExample) {
        setOutput(matchedExample.output);
      } else {
        const lines = [];

        // 2. Try to parse "jotokkhon" loop pattern (Simple Mock)
        // Pattern: jotokkhon (i < LIMIT) { dekho("STR" + i); i = i + 1; }
        const whileLoopMatch = code.match(/jotokkhon\s*\(\s*(\w+)\s*<\s*(\d+)\s*\)\s*\{/);
        const printInLoopMatch = code.match(/dekho\("([^"]+)"\s*\+\s*(\w+)\);/);

        if (whileLoopMatch && printInLoopMatch) {
          const varName = whileLoopMatch[1];
          const limit = parseInt(whileLoopMatch[2]);
          const prefix = printInLoopMatch[1];
          const printVar = printInLoopMatch[2];

          // Verify variable matches
          if (varName === printVar) {
            for (let i = 0; i < limit; i++) {
              lines.push(`${prefix}${i}`);
            }
          }
        }

        // 3. Try to parse simple print statements not in loop
        if (lines.length === 0) {
          const printMatches = code.matchAll(/dekho\("([^"]+)"\);/g);
          for (const match of printMatches) {
            lines.push(match[1]);
          }
        }

        // 4. Fallback if nothing matched
        if (lines.length === 0) {
          lines.push("Unrecognized code or custom logic.");
          lines.push("Note: This playground currently runs in simulation mode.");
          lines.push("Try the examples or specific patterns like:");
          lines.push('jotokkhon (i < 5) { dekho("Count: " + i); i = i + 1; }');
        }
        setOutput(lines);
      }

      setIsRunning(false);
    }, 600);
  };

  return (
    <div className="flex flex-col h-[calc(100vh-4rem)]">
      <div className="flex items-center justify-between px-6 py-3 border-b border-border bg-muted/20">
        <h1 className="text-lg font-semibold flex items-center gap-2">
          <span className="w-3 h-3 rounded-full bg-green-500"></span>
          Playground
        </h1>

        {/* Example Selector */}
        <div className="flex items-center gap-4">
          <div className="relative">
            <select
              value={selectedExample}
              onChange={handleExampleChange}
              className="appearance-none bg-background border border-border rounded-md px-4 py-1.5 pr-8 text-sm focus:outline-none focus:ring-1 focus:ring-primary cursor-pointer hover:border-primary/50 transition-colors"
            >
              {Object.keys(EXAMPLES).map(name => (
                <option key={name} value={name}>{name}</option>
              ))}
            </select>
            <ChevronDown className="absolute right-2 top-1/2 -translate-y-1/2 w-4 h-4 text-muted-foreground pointer-events-none" />
          </div>

          <div className="flex gap-2">
            <button
              onClick={() => {
                setCode(EXAMPLES[selectedExample as keyof typeof EXAMPLES].code);
                setOutput([]);
              }}
              className="flex items-center gap-2 px-4 py-2 text-sm font-medium text-muted-foreground hover:text-foreground transition-colors"
              title="Reset to original example code"
            >
              <RotateCcw className="w-4 h-4" /> Reset
            </button>
            <button
              onClick={handleRun}
              disabled={isRunning}
              className="flex items-center gap-2 px-6 py-2 text-sm font-medium text-white bg-green-600 hover:bg-green-700 rounded-md transition-colors disabled:opacity-50 shadow-lg shadow-green-900/20"
            >
              <Play className="w-4 h-4" /> {isRunning ? "Running..." : "Run Code"}
            </button>
          </div>
        </div>
      </div>

      <div className="flex-1 flex flex-col md:flex-row overflow-hidden">
        {/* Editor Pane */}
        <div className="flex-1 border-r border-border flex flex-col min-h-[50vh]">
          <div className="px-4 py-2 bg-muted/10 text-xs font-mono text-muted-foreground border-b border-border flex justify-between">
            <span>{selectedExample}</span>
            <span>BanglaCode Source</span>
          </div>
          <textarea
            value={code}
            onChange={(e) => setCode(e.target.value)}
            className="flex-1 w-full p-4 bg-[#1e1e1e] text-gray-300 font-mono text-sm leading-6 resize-none focus:outline-none"
            spellCheck={false}
          />
        </div>

        {/* Output Pane */}
        <div className="flex-1 flex flex-col bg-[#1e1e1e] min-h-[50vh]">
          <div className="px-4 py-2 bg-muted/10 text-xs font-mono text-muted-foreground border-b border-border flex justify-between">
            <span>Output</span>
            <span className="flex items-center gap-2 text-green-500/50">
              <span className="w-2 h-2 rounded-full bg-green-500 animate-pulse"></span>
              Connected
            </span>
          </div>
          <div className="flex-1 p-4 font-mono text-sm overflow-auto">
            {output.length > 0 ? (
              output.map((line, i) => (
                <div key={i} className="mb-1 text-gray-300 border-l-2 border-green-500/20 pl-3">
                  <span className="text-gray-600 mr-2 select-none">$</span>
                  {line}
                </div>
              ))
            ) : (
              <div className="text-gray-600 italic mt-8 text-center opacity-50">
                Ready to compile...
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
