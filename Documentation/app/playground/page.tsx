"use client";

import { useState, useCallback } from "react";
import { Play, RotateCcw, ChevronDown, AlertCircle, CheckCircle } from "lucide-react";
import { EXAMPLES } from "./examples";
import { execute } from "@/lib/interpreter";

export default function Playground() {
  const [selectedExample, setSelectedExample] = useState("hello.bang");
  const [code, setCode] = useState(EXAMPLES["hello.bang"].code);
  const [output, setOutput] = useState<string[]>([]);
  const [errors, setErrors] = useState<string[]>([]);
  const [isRunning, setIsRunning] = useState(false);
  const [executionTime, setExecutionTime] = useState<number | null>(null);

  const handleExampleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const key = e.target.value;
    setSelectedExample(key);
    setCode(EXAMPLES[key as keyof typeof EXAMPLES].code);
    setOutput([]);
    setErrors([]);
    setExecutionTime(null);
  };

  const handleRun = useCallback(() => {
    setIsRunning(true);
    setOutput([]);
    setErrors([]);
    setExecutionTime(null);

    // Small delay to show running state
    setTimeout(() => {
      const startTime = performance.now();
      
      // Execute the code using the real interpreter
      const result = execute(code);
      
      const endTime = performance.now();
      setExecutionTime(Math.round((endTime - startTime) * 100) / 100);

      setOutput(result.output);
      setErrors(result.errors);
      setIsRunning(false);
    }, 50);
  }, [code]);

  // Handle keyboard shortcut
  const handleKeyDown = useCallback((e: React.KeyboardEvent) => {
    if ((e.ctrlKey || e.metaKey) && e.key === "Enter") {
      e.preventDefault();
      handleRun();
    }
  }, [handleRun]);

  const examplesList = Object.entries(EXAMPLES).map(([key, value]) => ({
    key,
    name: value.name,
    description: value.description
  }));

  return (
    <div className="flex flex-col h-[calc(100vh-4rem)]">
      {/* Header */}
      <div className="flex flex-wrap items-center justify-between px-4 sm:px-6 py-3 border-b border-border bg-muted/20 gap-4">
        <div className="flex items-center gap-4">
          <h1 className="text-lg font-semibold flex items-center gap-2">
            <span className="w-3 h-3 rounded-full bg-green-500"></span>
            Playground
          </h1>
          <span className="text-xs text-muted-foreground hidden sm:inline">
            Ctrl+Enter to run
          </span>
        </div>

        {/* Example Selector */}
        <div className="flex items-center gap-2 sm:gap-4 flex-wrap">
          <div className="relative">
            <select
              value={selectedExample}
              onChange={handleExampleChange}
              className="appearance-none bg-background border border-border rounded-md px-3 sm:px-4 py-1.5 pr-8 text-sm focus:outline-none focus:ring-1 focus:ring-primary cursor-pointer hover:border-primary/50 transition-colors"
            >
              {examplesList.map(({ key, name }) => (
                <option key={key} value={key}>{name}</option>
              ))}
            </select>
            <ChevronDown className="absolute right-2 top-1/2 -translate-y-1/2 w-4 h-4 text-muted-foreground pointer-events-none" />
          </div>

          <div className="flex gap-2">
            <button
              onClick={() => {
                setCode(EXAMPLES[selectedExample as keyof typeof EXAMPLES].code);
                setOutput([]);
                setErrors([]);
                setExecutionTime(null);
              }}
              className="flex items-center gap-1 sm:gap-2 px-3 sm:px-4 py-2 text-sm font-medium text-muted-foreground hover:text-foreground transition-colors"
              title="Reset to original example code"
            >
              <RotateCcw className="w-4 h-4" />
              <span className="hidden sm:inline">Reset</span>
            </button>
            <button
              onClick={handleRun}
              disabled={isRunning}
              className="flex items-center gap-1 sm:gap-2 px-4 sm:px-6 py-2 text-sm font-medium text-white bg-green-600 hover:bg-green-700 rounded-md transition-colors disabled:opacity-50 shadow-lg shadow-green-900/20"
            >
              <Play className="w-4 h-4" />
              {isRunning ? "Running..." : "Run"}
            </button>
          </div>
        </div>
      </div>

      {/* Example description */}
      {EXAMPLES[selectedExample as keyof typeof EXAMPLES] && (
        <div className="px-4 sm:px-6 py-2 bg-primary/5 border-b border-border text-sm text-muted-foreground">
          <span className="font-medium text-foreground">{EXAMPLES[selectedExample as keyof typeof EXAMPLES].name}:</span>{" "}
          {EXAMPLES[selectedExample as keyof typeof EXAMPLES].description}
        </div>
      )}

      <div className="flex-1 flex flex-col lg:flex-row overflow-hidden">
        {/* Editor Pane */}
        <div className="flex-1 border-r border-border flex flex-col min-h-[40vh] lg:min-h-0">
          <div className="px-4 py-2 bg-muted/10 text-xs font-mono text-muted-foreground border-b border-border flex justify-between">
            <span>{selectedExample}</span>
            <span>BanglaCode Source</span>
          </div>
          <textarea
            value={code}
            onChange={(e) => setCode(e.target.value)}
            onKeyDown={handleKeyDown}
            className="flex-1 w-full p-4 bg-[#1e1e1e] text-gray-300 font-mono text-sm leading-6 resize-none focus:outline-none"
            spellCheck={false}
            placeholder="Write your BanglaCode here..."
          />
        </div>

        {/* Output Pane */}
        <div className="flex-1 flex flex-col bg-[#1e1e1e] min-h-[40vh] lg:min-h-0">
          <div className="px-4 py-2 bg-muted/10 text-xs font-mono text-muted-foreground border-b border-border flex justify-between items-center">
            <span>Output</span>
            <div className="flex items-center gap-4">
              {executionTime !== null && (
                <span className="text-muted-foreground">
                  {executionTime}ms
                </span>
              )}
              {errors.length > 0 ? (
                <span className="flex items-center gap-1 text-red-400">
                  <AlertCircle className="w-3 h-3" />
                  Error
                </span>
              ) : output.length > 0 ? (
                <span className="flex items-center gap-1 text-green-400">
                  <CheckCircle className="w-3 h-3" />
                  Success
                </span>
              ) : (
                <span className="flex items-center gap-2 text-green-500/50">
                  <span className="w-2 h-2 rounded-full bg-green-500 animate-pulse"></span>
                  Ready
                </span>
              )}
            </div>
          </div>
          <div className="flex-1 p-4 font-mono text-sm overflow-auto">
            {errors.length > 0 && (
              <div className="mb-4">
                {errors.map((error, i) => (
                  <div key={`error-${i}`} className="mb-1 text-red-400 border-l-2 border-red-500/50 pl-3">
                    <span className="text-red-600 mr-2 select-none">✗</span>
                    {error}
                  </div>
                ))}
              </div>
            )}
            {output.length > 0 ? (
              output.map((line, i) => (
                <div key={i} className="mb-1 text-gray-300 border-l-2 border-green-500/20 pl-3">
                  <span className="text-gray-600 mr-2 select-none">$</span>
                  {line}
                </div>
              ))
            ) : errors.length === 0 ? (
              <div className="text-gray-600 italic mt-8 text-center opacity-50">
                <p>Press <kbd className="px-2 py-1 bg-gray-800 rounded text-gray-400">Run</kbd> or <kbd className="px-2 py-1 bg-gray-800 rounded text-gray-400">Ctrl+Enter</kbd> to execute</p>
                <p className="mt-2 text-xs">The code runs entirely in your browser - no server needed!</p>
              </div>
            ) : null}
          </div>
        </div>
      </div>

      {/* Footer */}
      <div className="px-4 sm:px-6 py-2 border-t border-border bg-muted/10 text-xs text-muted-foreground flex flex-wrap justify-between gap-2">
        <span>BanglaCode Playground - Write and run code in your browser</span>
        <span className="hidden sm:inline">
          {Object.keys(EXAMPLES).length} examples available
        </span>
      </div>
    </div>
  );
}
