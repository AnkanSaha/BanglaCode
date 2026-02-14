import Link from "next/link";
import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";
import { Download, Terminal, Box, Apple, Monitor, Server } from "lucide-react";

export default function Installation() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Getting Started
        </span>
      </div>

      <h1>Installation</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        Install BanglaCode on your machine and start coding in Bengali.
        Use our one-liner installers, download pre-built binaries, or build from source.
      </p>

      <h2>🚀 Quick Install (Recommended)</h2>

      <p>
        The fastest way to install BanglaCode. Our installer automatically detects your
        operating system and architecture.
      </p>

      <h3>Linux / macOS / BSD</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`curl -fsSL https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.sh | bash`}
      />

      <p className="text-sm text-muted-foreground mt-2">
        Supports: Ubuntu, Debian, Fedora, RHEL, CentOS, Arch, macOS (Intel &amp; Apple Silicon), 
        FreeBSD, OpenBSD, NetBSD, and more.
      </p>

      <h3>Windows (PowerShell)</h3>

      <CodeBlock
        language="powershell"
        showLineNumbers={false}
        code={`irm https://raw.githubusercontent.com/nexoral/BanglaCode/main/Scripts/install.ps1 | iex`}
      />

      <p className="text-sm text-muted-foreground mt-2">
        Works on Windows 7+, Windows 10/11, and Windows Server. Supports x64, x86, and ARM64.
      </p>

      <h2>📦 Download Pre-built Binaries</h2>

      <p>
        Download the appropriate package for your system from our{" "}
        <a
          href="https://github.com/nexoral/BanglaCode/releases"
          target="_blank"
          rel="noopener noreferrer"
          className="text-primary hover:underline"
        >
          GitHub Releases
        </a>.
      </p>

      <div className="grid gap-6 my-6">
        {/* Windows */}
        <div className="p-4 rounded-lg border border-border bg-secondary/10">
          <div className="flex items-center gap-2 mb-3">
            <Monitor className="w-5 h-5 text-blue-500" />
            <h4 className="font-medium m-0">Windows</h4>
          </div>
          <div className="overflow-x-auto">
            <table className="text-sm w-full">
              <thead>
                <tr className="border-b border-border">
                  <th className="text-left py-2 pr-4">File</th>
                  <th className="text-left py-2 pr-4">Architecture</th>
                  <th className="text-left py-2">Description</th>
                </tr>
              </thead>
              <tbody>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-windows-amd64.zip</td>
                  <td className="py-2 pr-4">x64</td>
                  <td className="py-2 text-muted-foreground">64-bit (most common)</td>
                </tr>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-windows-386.zip</td>
                  <td className="py-2 pr-4">x86</td>
                  <td className="py-2 text-muted-foreground">32-bit</td>
                </tr>
                <tr>
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-windows-arm64.zip</td>
                  <td className="py-2 pr-4">ARM64</td>
                  <td className="py-2 text-muted-foreground">Surface Pro X, Windows on ARM</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        {/* macOS */}
        <div className="p-4 rounded-lg border border-border bg-secondary/10">
          <div className="flex items-center gap-2 mb-3">
            <Apple className="w-5 h-5 text-gray-500" />
            <h4 className="font-medium m-0">macOS</h4>
          </div>
          <div className="overflow-x-auto">
            <table className="text-sm w-full">
              <thead>
                <tr className="border-b border-border">
                  <th className="text-left py-2 pr-4">File</th>
                  <th className="text-left py-2 pr-4">Architecture</th>
                  <th className="text-left py-2">Description</th>
                </tr>
              </thead>
              <tbody>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-macos-arm64.tar.gz</td>
                  <td className="py-2 pr-4">ARM64</td>
                  <td className="py-2 text-muted-foreground">Apple Silicon (M1/M2/M3) ⭐</td>
                </tr>
                <tr>
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-macos-amd64.tar.gz</td>
                  <td className="py-2 pr-4">x64</td>
                  <td className="py-2 text-muted-foreground">Intel Mac</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        {/* Linux */}
        <div className="p-4 rounded-lg border border-border bg-secondary/10">
          <div className="flex items-center gap-2 mb-3">
            <Server className="w-5 h-5 text-orange-500" />
            <h4 className="font-medium m-0">Linux</h4>
          </div>
          <div className="overflow-x-auto">
            <table className="text-sm w-full">
              <thead>
                <tr className="border-b border-border">
                  <th className="text-left py-2 pr-4">File</th>
                  <th className="text-left py-2 pr-4">Architecture</th>
                  <th className="text-left py-2">Description</th>
                </tr>
              </thead>
              <tbody>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-linux-amd64.deb</td>
                  <td className="py-2 pr-4">x64</td>
                  <td className="py-2 text-muted-foreground">Debian/Ubuntu</td>
                </tr>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-linux-amd64.rpm</td>
                  <td className="py-2 pr-4">x64</td>
                  <td className="py-2 text-muted-foreground">Fedora/RHEL/CentOS</td>
                </tr>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-linux-amd64.tar.gz</td>
                  <td className="py-2 pr-4">x64</td>
                  <td className="py-2 text-muted-foreground">Generic Linux</td>
                </tr>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-linux-arm64.deb</td>
                  <td className="py-2 pr-4">ARM64</td>
                  <td className="py-2 text-muted-foreground">Raspberry Pi 4+, AWS Graviton</td>
                </tr>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-linux-arm64.rpm</td>
                  <td className="py-2 pr-4">ARM64</td>
                  <td className="py-2 text-muted-foreground">ARM64 Fedora/RHEL</td>
                </tr>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-linux-arm.tar.gz</td>
                  <td className="py-2 pr-4">ARMv7</td>
                  <td className="py-2 text-muted-foreground">Raspberry Pi 2/3, embedded</td>
                </tr>
                <tr>
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-linux-386.tar.gz</td>
                  <td className="py-2 pr-4">x86</td>
                  <td className="py-2 text-muted-foreground">32-bit Linux</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        {/* BSD */}
        <div className="p-4 rounded-lg border border-border bg-secondary/10">
          <div className="flex items-center gap-2 mb-3">
            <Terminal className="w-5 h-5 text-red-500" />
            <h4 className="font-medium m-0">BSD</h4>
          </div>
          <div className="overflow-x-auto">
            <table className="text-sm w-full">
              <thead>
                <tr className="border-b border-border">
                  <th className="text-left py-2 pr-4">File</th>
                  <th className="text-left py-2 pr-4">Architecture</th>
                  <th className="text-left py-2">Description</th>
                </tr>
              </thead>
              <tbody>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-freebsd-amd64.tar.gz</td>
                  <td className="py-2 pr-4">x64</td>
                  <td className="py-2 text-muted-foreground">FreeBSD 64-bit</td>
                </tr>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-freebsd-386.tar.gz</td>
                  <td className="py-2 pr-4">x86</td>
                  <td className="py-2 text-muted-foreground">FreeBSD 32-bit</td>
                </tr>
                <tr className="border-b border-border/50">
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-openbsd-amd64.tar.gz</td>
                  <td className="py-2 pr-4">x64</td>
                  <td className="py-2 text-muted-foreground">OpenBSD 64-bit</td>
                </tr>
                <tr>
                  <td className="py-2 pr-4 font-mono text-xs">banglacode-netbsd-amd64.tar.gz</td>
                  <td className="py-2 pr-4">x64</td>
                  <td className="py-2 text-muted-foreground">NetBSD 64-bit</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <h2>📋 Manual Installation</h2>

      <h3>Linux (DEB - Debian/Ubuntu)</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Download the .deb file, then:
sudo dpkg -i banglacode-linux-amd64.deb`}
      />

      <h3>Linux (RPM - Fedora/RHEL/CentOS)</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# With dnf
sudo dnf install ./banglacode-linux-amd64.rpm

# Or with yum
sudo yum install ./banglacode-linux-amd64.rpm`}
      />

      <h3>Linux/macOS/BSD (tar.gz)</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Extract and install
tar -xzf banglacode-<os>-<arch>.tar.gz
sudo mv banglacode /usr/local/bin/
sudo chmod +x /usr/local/bin/banglacode`}
      />

      <h3>Windows (ZIP)</h3>

      <ol>
        <li>Download the appropriate ZIP file</li>
        <li>Extract to a folder (e.g., <code>C:\Program Files\BanglaCode</code>)</li>
        <li>Add the folder to your system PATH:
          <ul>
            <li>Search &quot;Environment Variables&quot; in Windows</li>
            <li>Edit &quot;Path&quot; under User variables</li>
            <li>Add the folder containing <code>banglacode.exe</code></li>
          </ul>
        </li>
        <li>Restart your terminal</li>
      </ol>

      <h2>🔐 Verify Download (Optional)</h2>

      <p>
        Each release includes a <code>checksums.txt</code> file with SHA256 hashes.
        Verify your download:
      </p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Download checksums.txt from the release, then:
sha256sum -c checksums.txt --ignore-missing`}
      />

      <h2>🔧 Build from Source</h2>

      <h3>Prerequisites</h3>

      <p>
        To build from source, you need <strong>Go 1.20 or later</strong>.
        Download it from <a href="https://go.dev/dl/" target="_blank" rel="noopener noreferrer" className="text-primary hover:underline">go.dev/dl</a>.
      </p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`go version
# Should show: go version go1.20+ ...`}
      />

      <h3>Build Steps</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Clone the repository
git clone https://github.com/nexoral/BanglaCode.git
cd BanglaCode

# Build
go build -o banglacode main.go

# Verify
./banglacode --version

# Optional: Install system-wide
sudo mv banglacode /usr/local/bin/`}
      />

      <h3>Cross-Compilation</h3>

      <p>Build for different platforms:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -o banglacode.exe main.go

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o banglacode main.go

# Linux (ARM64 - Raspberry Pi 4)
GOOS=linux GOARCH=arm64 go build -o banglacode main.go

# FreeBSD
GOOS=freebsd GOARCH=amd64 go build -o banglacode main.go`}
      />

      <h2>✅ Verify Installation</h2>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`banglacode --version`}
      />

      <h2>🎯 Running Your First Program</h2>

      <h3>Create a File</h3>

      <p>Create a file named <code>hello.bang</code>:</p>

      <CodeBlock
        filename="hello.bang"
        code={`// Your first BanglaCode program
dekho("Namaskar, BanglaCode!");

dhoro naam = "World";
dekho("Hello,", naam);`}
      />

      <h3>Run the Program</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`banglacode hello.bang`}
      />

      <p>Output:</p>

      <CodeBlock
        language="output"
        showLineNumbers={false}
        code={`Namaskar, BanglaCode!
Hello, World`}
      />

      <h2>💻 Interactive REPL</h2>

      <p>Start the interactive mode by running without arguments:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`banglacode`}
      />

      <p>You&apos;ll see a prompt where you can type BanglaCode directly:</p>

      <CodeBlock
        language="output"
        showLineNumbers={false}
        code={`BanglaCode REPL v1.0
Type 'help' for commands, 'exit' to quit

>> dhoro x = 5
>> dekho(x * 2)
10
>> kaj double(n) { ferao n * 2; }
>> double(21)
42`}
      />

      <h2>🎨 VS Code Extension</h2>

      <p>
        For syntax highlighting and code snippets, install the BanglaCode VS Code extension:
      </p>

      <ol>
        <li>Open VS Code</li>
        <li>Go to Extensions (Ctrl+Shift+X)</li>
        <li>Search for &quot;BanglaCode&quot;</li>
        <li>Click Install</li>
      </ol>

      <p>
        Or visit the{" "}
        <a
          href="https://marketplace.visualstudio.com/items?itemName=AnkanSaha.banglacode"
          target="_blank"
          rel="noopener noreferrer"
          className="text-primary hover:underline"
        >
          VS Code Marketplace
        </a>.
      </p>

      <h2>❓ Troubleshooting</h2>

      <h3>Command not found</h3>

      <p>Make sure the binary is in your PATH:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`# Check if banglacode is in PATH
which banglacode

# If not found, add to PATH (Linux/macOS)
echo 'export PATH="$PATH:/usr/local/bin"' >> ~/.bashrc
source ~/.bashrc`}
      />

      <h3>Permission denied (Linux/macOS)</h3>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`chmod +x /usr/local/bin/banglacode`}
      />

      <h3>Windows: Script execution disabled</h3>

      <p>If the PowerShell installer is blocked:</p>

      <CodeBlock
        language="powershell"
        showLineNumbers={false}
        code={`Set-ExecutionPolicy RemoteSigned -Scope CurrentUser`}
      />

      <h3>Build errors</h3>

      <p>Make sure you have Go 1.20+ and all dependencies:</p>

      <CodeBlock
        language="bash"
        showLineNumbers={false}
        code={`go mod tidy
go build -o banglacode main.go`}
      />

      <DocNavigation currentPath="/docs/installation" />
    </div>
  );
}
