"use client";

import { useState } from "react";
import { Play, RotateCcw } from "lucide-react";

const EXAMPLE_CODE = `// Hello World in BanglaCode
kaj greet(name) {
  dekho("Namaskar, " + name);
}

greet("Bangla");

// Loop example
dhoro i = 0;
jotokkhon (i < 5) {
  dekho("Count: " + i);
  i = i + 1;
}`;

export default function Playground() {
  const [code, setCode] = useState(EXAMPLE_CODE);
  const [output, setOutput] = useState<string[]>([]);
  const [isRunning, setIsRunning] = useState(false);

  const handleRun = async () => {
    setIsRunning(true);
    setOutput([]);

    // Simulating compilation delay
    setTimeout(() => {
      // Mock output parser for demo
      const lines = [];
      if (code.includes('dekho("Namaskar, " + name)')) {
        lines.push("Namaskar, Bangla");
      }
      if (code.includes('dekho("Count: " + i)')) {
        for (let i = 0; i < 5; i++) lines.push(`Count: ${i}`);
      }

      if (lines.length === 0) {
        lines.push("Output will appear here when you run the code...");
        lines.push("(Note: This is a playground demo. Actual WASM compilation coming soon!)");
      }

      setOutput(lines);
      setIsRunning(false);
    }, 800);
  };

  return (
    <div className="flex flex-col h-[calc(100vh-4rem)]">
      <div className="flex items-center justify-between px-6 py-3 border-b border-border bg-muted/20">
        <h1 className="text-lg font-semibold flex items-center gap-2">
          <span className="w-3 h-3 rounded-full bg-green-500"></span>
          Playground
        </h1>
        <div className="flex gap-2">
          <button
            onClick={() => setCode(EXAMPLE_CODE)}
            className="flex items-center gap-2 px-4 py-2 text-sm font-medium text-muted-foreground hover:text-foreground transition-colors"
          >
            <RotateCcw className="w-4 h-4" /> Reset
          </button>
          <button
            onClick={handleRun}
            disabled={isRunning}
            className="flex items-center gap-2 px-6 py-2 text-sm font-medium text-white bg-green-600 hover:bg-green-700 rounded-md transition-colors disabled:opacity-50"
          >
            <Play className="w-4 h-4" /> {isRunning ? "Running..." : "Run Code"}
          </button>
        </div>
      </div>

      <div className="flex-1 flex flex-col md:flex-row overflow-hidden">
        {/* Editor Pane */}
        <div className="flex-1 border-r border-border flex flex-col min-h-[50vh]">
          <div className="px-4 py-2 bg-muted/10 text-xs font-mono text-muted-foreground border-b border-border">main.bang</div>
          <textarea
            value={code}
            onChange={(e) => setCode(e.target.value)}
            className="flex-1 w-full p-4 bg-[#1e1e1e] text-gray-300 font-mono text-sm resize-none focus:outline-none"
            spellCheck={false}
          />
        </div>

        {/* Output Pane */}
        <div className="flex-1 flex flex-col bg-[#1e1e1e] min-h-[50vh]">
          <div className="px-4 py-2 bg-muted/10 text-xs font-mono text-muted-foreground border-b border-border flex justify-between">
            <span>Output</span>
            <span className="text-green-500/50">● Connected</span>
          </div>
          <div className="flex-1 p-4 font-mono text-sm overflow-auto">
            {output.length > 0 ? (
              output.map((line, i) => (
                <div key={i} className="mb-1 text-gray-300">
                  <span className="text-gray-600 mr-2">$</span>
                  {line}
                </div>
              ))
            ) : (
              <div className="text-gray-600 italic">
                Ready to compile...
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
