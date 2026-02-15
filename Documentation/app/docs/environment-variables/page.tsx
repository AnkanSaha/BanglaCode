import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function EnvironmentVariables() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
      </div>

      <h1>Environment Variables</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides built-in support for loading and managing environment variables from <code>.env</code> files,
        with multi-environment support for development, staging, and production configurations.
      </p>

      <h2>Why Use .env Files?</h2>

      <p>
        Environment variables help you manage configuration settings across different environments without hardcoding
        sensitive data like API keys, database passwords, or URLs in your code.
      </p>

      <ul>
        <li><strong>Security</strong> - Keep secrets out of version control</li>
        <li><strong>Flexibility</strong> - Different configs for dev, staging, production</li>
        <li><strong>Portability</strong> - Same code works across environments</li>
        <li><strong>Best Practice</strong> - Industry-standard configuration management</li>
      </ul>

      <h2>Environment Variable Functions</h2>

      <div className="grid gap-4 my-6">
        <div className="border rounded-lg p-4">
          <h3 className="mt-0"><code>env_load(filename)</code></h3>
          <p className="text-muted-foreground">Load environment variables from a specific .env file</p>
        </div>

        <div className="border rounded-lg p-4">
          <h3 className="mt-0"><code>env_load_auto(environment)</code></h3>
          <p className="text-muted-foreground">
            Automatically load <code>.env.{'{environment}'}</code> or fallback to <code>.env</code>
          </p>
        </div>

        <div className="border rounded-lg p-4">
          <h3 className="mt-0"><code>env_get(key)</code></h3>
          <p className="text-muted-foreground">Get environment variable value (throws error if not found)</p>
        </div>

        <div className="border rounded-lg p-4">
          <h3 className="mt-0"><code>env_get_default(key, default)</code></h3>
          <p className="text-muted-foreground">Get environment variable with a default fallback value</p>
        </div>

        <div className="border rounded-lg p-4">
          <h3 className="mt-0"><code>env_set(key, value)</code></h3>
          <p className="text-muted-foreground">Set environment variable at runtime</p>
        </div>

        <div className="border rounded-lg p-4">
          <h3 className="mt-0"><code>env_all()</code></h3>
          <p className="text-muted-foreground">Get all environment variables as a map</p>
        </div>

        <div className="border rounded-lg p-4">
          <h3 className="mt-0"><code>env_clear()</code></h3>
          <p className="text-muted-foreground">Clear all loaded environment variables</p>
        </div>
      </div>

      <h2>Creating .env Files</h2>

      <p>
        Create a <code>.env</code> file in your project root with <code>KEY=VALUE</code> pairs:
      </p>

      <CodeBlock
        language="bash"
        code={`# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=myapp
DB_USER=admin
DB_PASSWORD=secret123

# API Configuration
API_KEY=your_api_key_here
API_URL=http://localhost:3000
SECRET_KEY=your_secret_key

# Application Settings
APP_NAME=My BanglaCode App
NODE_ENV=development
DEBUG=true
PORT=8080`}
      />

      <h2>Loading Environment Variables</h2>

      <h3>Basic Loading</h3>

      <CodeBlock
        code={`// Load .env file
env_load(".env");

// Get environment variables
dhoro api_key = env_get("API_KEY");
dhoro api_url = env_get("API_URL");
dhoro app_name = env_get("APP_NAME");

dekho("App Name:", app_name);
dekho("API URL:", api_url);
dekho("API Key:", api_key);`}
      />

      <h3>Using Default Values</h3>

      <CodeBlock
        code={`// Load with defaults (safer)
env_load(".env");

dhoro api_url = env_get_default("API_URL", "http://localhost:3000");
dhoro debug = env_get_default("DEBUG", "false");
dhoro port = env_get_default("PORT", "8080");

dekho("API URL:", api_url);
dekho("Debug Mode:", debug);
dekho("Server Port:", port);`}
      />

      <h2>Multi-Environment Support</h2>

      <p>
        BanglaCode supports multiple environment files: <code>.env</code>, <code>.env.dev</code>,
        <code>.env.uat</code>, <code>.env.staging</code>, <code>.env.prod</code>
      </p>

      <h3>Environment-Specific Files</h3>

      <div className="my-4">
        <p><strong>.env</strong> (Default/Development)</p>
        <CodeBlock
          language="bash"
          code={`API_URL=http://localhost:3000
DB_HOST=localhost
DEBUG=true`}
        />

        <p className="mt-4"><strong>.env.prod</strong> (Production)</p>
        <CodeBlock
          language="bash"
          code={`API_URL=https://api.production.com
DB_HOST=prod-db.example.com
DEBUG=false`}
        />

        <p className="mt-4"><strong>.env.uat</strong> (User Acceptance Testing)</p>
        <CodeBlock
          language="bash"
          code={`API_URL=https://api.uat.example.com
DB_HOST=uat-db.example.com
DEBUG=true`}
        />
      </div>

      <h3>Auto-Loading with Fallback</h3>

      <CodeBlock
        code={`// Load production environment
// Tries .env.prod first, then falls back to .env if not found
env_load_auto("prod");

dhoro api_url = env_get("API_URL");
dekho("Production API:", api_url);  // Uses .env.prod value

// Load UAT environment
env_clear();  // Clear previous env vars
env_load_auto("uat");

dhoro uat_url = env_get("API_URL");
dekho("UAT API:", uat_url);  // Uses .env.uat value

// Load non-existent environment (fallbacks to .env)
env_clear();
env_load_auto("staging");  // No .env.staging, uses .env

dhoro default_url = env_get("API_URL");
dekho("Fallback API:", default_url);  // Uses .env value`}
      />

      <h2>Runtime Environment Variables</h2>

      <CodeBlock
        code={`// Load .env file
env_load(".env");

// Set runtime variables
env_set("CURRENT_USER", "Rahim Ahmed");
env_set("SESSION_ID", "abc123xyz");
env_set("REQUEST_COUNT", "42");

// Get runtime variables
dhoro user = env_get("CURRENT_USER");
dhoro session = env_get("SESSION_ID");

dekho("Current User:", user);
dekho("Session ID:", session);`}
      />

      <h2>Practical Examples</h2>

      <h3>Database Connection with .env</h3>

      <CodeBlock
        code={`// Load environment-specific config
env_load_auto("prod");

// Get database credentials from .env
dhoro db_host = env_get("DB_HOST");
dhoro db_port = env_get_default("DB_PORT", "5432");
dhoro db_name = env_get("DB_NAME");
dhoro db_user = env_get("DB_USER");
dhoro db_pass = env_get("DB_PASSWORD");

// Connect to database using env variables
dhoro conn = db_jukto("postgres", {
    "host": db_host,
    "port": sonkha(db_port),
    "database": db_name,
    "user": db_user,
    "password": db_pass
});

dhoro users = db_query(conn, "SELECT * FROM users");
dekho("Found", dorghyo(users["rows"]), "users");

db_bandho(conn);`}
      />

      <h3>HTTP Server with Configuration</h3>

      <CodeBlock
        code={`// Load environment config
env_load_auto("dev");

// Get server settings
dhoro host = env_get_default("HOST", "0.0.0.0");
dhoro port = sonkha(env_get_default("PORT", "8080"));
dhoro app_name = env_get_default("APP_NAME", "My App");

dekho("Starting", app_name, "on", host + ":" + lipi(port));

// Start server
server_chalu(port, kaj(req) {
    dhoro api_key = env_get("API_KEY");

    // Validate API key from request
    jodi (req.headers["x-api-key"] != api_key) {
        ferao json_uttor(401, {error: "Invalid API key"});
    }

    ferao json_uttor(200, {
        message: "Welcome to " + app_name,
        environment: env_get_default("NODE_ENV", "development")
    });
});`}
      />

      <h3>Multi-Environment Deployment Script</h3>

      <CodeBlock
        code={`kaj deploy(environment) {
    dekho("========================================");
    dekho("Deploying to:", boroHater(environment));
    dekho("========================================");

    // Load environment-specific config
    env_load_auto(environment);

    // Display configuration
    dekho("");
    dekho("Configuration:");
    dekho("  API URL:", env_get("API_URL"));
    dekho("  DB Host:", env_get("DB_HOST"));
    dekho("  Debug Mode:", env_get_default("DEBUG", "false"));
    dekho("");

    // Get all env vars for deployment
    dhoro all_vars = env_all();
    dekho("Total env variables:", dorghyo(chabi(all_vars)));

    // Deployment logic here...
    dekho("Deployment successful!");
}

// Usage
deploy("uat");    // Deploy to UAT
deploy("prod");   // Deploy to production`}
      />

      <h3>Configuration Manager</h3>

      <CodeBlock
        code={`sreni Config {
    shuru(environment) {
        ei.env = environment;
        ei.loaded = mittha;
        ei.load();
    }

    kaj load() {
        dekho("Loading config for:", ei.env);
        env_load_auto(ei.env);
        ei.loaded = sotti;
    }

    kaj get(key, defaultValue) {
        jodi (na ei.loaded) {
            ei.load();
        }

        jodi (defaultValue != khali) {
            ferao env_get_default(key, defaultValue);
        }

        ferao env_get(key);
    }

    kaj reload() {
        env_clear();
        ei.loaded = mittha;
        ei.load();
    }

    kaj displayAll() {
        dhoro vars = env_all();
        dhoro keys = chabi(vars);

        dekho("Environment Variables:");
        ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
            dhoro key = keys[i];
            dekho("  ", key, "=", vars[key]);
        }
    }
}

// Usage
dhoro config = notun Config("prod");

dhoro api_url = config.get("API_URL");
dhoro db_host = config.get("DB_HOST");
dhoro port = config.get("PORT", "3000");

dekho("API URL:", api_url);
dekho("DB Host:", db_host);
dekho("Port:", port);

// Display all config
config.displayAll();

// Reload config
config.reload();`}
      />

      <h3>Secret Validation</h3>

      <CodeBlock
        code={`kaj validateSecrets() {
    dhoro required = [
        "API_KEY",
        "DB_PASSWORD",
        "SECRET_KEY"
    ];

    dhoro missing = [];

    ghuriye (dhoro i = 0; i < dorghyo(required); i = i + 1) {
        dhoro key = required[i];

        chesta {
            env_get(key);
        } dhoro_bhul (error) {
            dhokao(missing, key);
        }
    }

    jodi (dorghyo(missing) > 0) {
        dekho("ERROR: Missing required environment variables:");
        ghuriye (dhoro i = 0; i < dorghyo(missing); i = i + 1) {
            dekho("  -", missing[i]);
        }
        ferao mittha;
    }

    dekho("All required secrets present!");
    ferao sotti;
}

// Load config
env_load_auto("prod");

// Validate before starting app
jodi (validateSecrets()) {
    dekho("Starting application...");
    // Start app logic
} nahole {
    dekho("Cannot start app without required secrets!");
}`}
      />

      <h2>Best Practices</h2>

      <ul>
        <li><strong>Never commit .env files</strong> - Add <code>.env*</code> to <code>.gitignore</code></li>
        <li><strong>Use .env.example</strong> - Commit a template without sensitive values</li>
        <li><strong>Use descriptive variable names</strong> - <code>DB_HOST</code> not <code>H1</code></li>
        <li><strong>Provide defaults</strong> - Use <code>env_get_default()</code> for non-critical values</li>
        <li><strong>Validate secrets</strong> - Check required variables exist before starting</li>
        <li><strong>Document environment files</strong> - Explain what each variable does</li>
        <li><strong>Use environment-specific files</strong> - <code>.env.dev</code>, <code>.env.prod</code></li>
        <li><strong>Rotate secrets regularly</strong> - Update API keys and passwords periodically</li>
      </ul>

      <h2>Security Tips</h2>

      <div className="bg-yellow-50 dark:bg-yellow-950 border border-yellow-200 dark:border-yellow-800 rounded-lg p-4 my-6">
        <h3 className="mt-0 text-yellow-900 dark:text-yellow-100">⚠️ Important Security Notes</h3>
        <ul className="mb-0 text-yellow-900 dark:text-yellow-100">
          <li><strong>Never log sensitive values</strong> - Don&apos;t use <code>dekho()</code> with passwords or API keys</li>
          <li><strong>Use .gitignore</strong> - Prevent accidentally committing <code>.env</code> files</li>
          <li><strong>Encrypt production secrets</strong> - Use secret management tools for production</li>
          <li><strong>Limit access</strong> - Only authorized people should have access to .env files</li>
          <li><strong>Monitor for leaks</strong> - Use tools to scan for exposed secrets</li>
        </ul>
      </div>

      <h2>Example .env.example Template</h2>

      <CodeBlock
        language="bash"
        code={`# Example .env file for BanglaCode
# Copy this to .env and fill in your actual values

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_NAME=myapp
DB_USER=admin
DB_PASSWORD=your_password_here

# API Configuration
API_KEY=your_api_key_here
API_URL=http://localhost:3000
SECRET_KEY=your_secret_key_here

# Application Settings
APP_NAME=BanglaCode App
NODE_ENV=development
DEBUG=true
PORT=8080

# Email Settings (optional)
EMAIL_HOST=smtp.gmail.com
EMAIL_PORT=587
EMAIL_USER=your_email@gmail.com
EMAIL_PASSWORD=your_email_password`}
      />

      <DocNavigation currentPath="/docs/environment-variables" />
    </div>
  );
}
