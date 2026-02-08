import Link from "next/link";
import { ArrowRight, Terminal } from "lucide-react";

export default function Installation() {
  return (
    <div className="space-y-6">
      <h1>Installation</h1>
      <p className="lead text-xl text-muted-foreground">
        Get started with BanglaCode on your machine.
      </p>

      <h2>Prerequisites</h2>
      <p>
        Currently, BanglaCode is built using Go. You will need <strong>Go 1.20+</strong> installed on your system if you want to build from source.
        <br />
        <Link href="https://github.com/nexoral/BanglaCode/releases" className="text-primary hover:underline font-medium" target="_blank">
          Download Pre-built Binaries from GitHub
        </Link>
      </p>

      <h2>Building from Source</h2>

      <div className="step space-y-4">
        <h3 className="flex items-center gap-2">
          <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary/20 text-primary text-sm font-bold">1</span>
          Clone the Repository
        </h3>
        <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
          git clone https://github.com/nexoral/BanglaCode.git<br />
          cd BanglaCode
        </div>
      </div>

      <div className="step space-y-4 pt-4">
        <h3 className="flex items-center gap-2">
          <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary/20 text-primary text-sm font-bold">2</span>
          Build the Binary
        </h3>
        <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
          go build -o banglacode main.go
        </div>
      </div>

      <div className="step space-y-4 pt-4">
        <h3 className="flex items-center gap-2">
          <span className="flex items-center justify-center w-8 h-8 rounded-full bg-primary/20 text-primary text-sm font-bold">3</span>
          Verify Installation
        </h3>
        <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
          ./banglacode --version
        </div>
      </div>

      <h2>Running Your First Code</h2>
      <p>Create a file named <code>hello.bang</code> and add the following:</p>

      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        dekho("Joy Bangla!");
      </div>

      <p>Run it utilizing the interpreter:</p>
      <div className="bg-secondary/50 p-4 rounded-lg font-mono text-sm border border-border">
        ./banglacode hello.bang
      </div>

      <div className="pt-8">
        <Link
          href="/docs/syntax"
          className="inline-flex items-center gap-2 px-6 py-3 bg-primary text-white rounded-lg hover:bg-primary/90 transition-colors"
        >
          Next: Syntax Guide <ArrowRight className="w-4 h-4" />
        </Link>
      </div>
    </div>
  );
}
