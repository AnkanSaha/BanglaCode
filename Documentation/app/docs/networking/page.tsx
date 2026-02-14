import Link from "next/link";
import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function Networking() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
      </div>

      <h1>Networking (TCP, UDP, WebSocket)</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode provides comprehensive networking capabilities for building servers and clients
        using TCP, UDP, and WebSocket protocols - as easy to use as JavaScript/Node.js.
      </p>

      <h2>TCP (Transmission Control Protocol)</h2>

      <p>
        TCP provides reliable, ordered, and error-checked delivery of data. Perfect for applications
        that need guaranteed message delivery like chat servers, file transfers, and API clients.
      </p>

      <h3>TCP Server</h3>

      <p>
        Use <code>tcp_server_chalu(port, handler)</code> to create a TCP server. The handler function
        receives a connection object for each client connection.
      </p>

      <CodeBlock
        filename="tcp_server.bang"
        code={`// TCP Echo Server
dekho("Starting TCP server on port 8080...");

kaj handleConnection(conn) {
    dekho("Client connected:", conn["remote_addr"]);

    // Echo back the received data
    dhoro data = conn["data"];
    tcp_pathao(conn, "Echo: " + data);

    dekho("Sent response");
}

tcp_server_chalu(8080, handleConnection);
dekho("Server running!");

// Keep server alive
jotokkhon (sotti) {
    ghumaao(1000);
}`}
      />

      <h3>TCP Client</h3>

      <p>
        Use <code>tcp_jukto(host, port)</code> to connect to a TCP server. This function is async
        and returns a promise.
      </p>

      <CodeBlock
        filename="tcp_client.bang"
        code={`proyash kaj tcpClient() {
    dekho("Connecting to TCP server...");

    // Connect to server (async)
    dhoro conn = opekha tcp_jukto("localhost", 8080);
    dekho("Connected!");

    // Send message
    tcp_lekho(conn, "Hello from BanglaCode!");

    // Read response (async)
    dhoro response = opekha tcp_shuno(conn);
    dekho("Server replied:", response);

    // Close connection
    tcp_bondho(conn);
}

tcpClient();`}
      />

      <h3>TCP Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>tcp_server_chalu</code></td>
              <td><code>port, handler</code></td>
              <td>Start TCP server with callback for each connection</td>
            </tr>
            <tr>
              <td><code>tcp_jukto</code></td>
              <td><code>host, port</code></td>
              <td>Connect to TCP server (async, returns promise)</td>
            </tr>
            <tr>
              <td><code>tcp_pathao</code></td>
              <td><code>connection, data</code></td>
              <td>Send data on TCP connection</td>
            </tr>
            <tr>
              <td><code>tcp_lekho</code></td>
              <td><code>connection, data</code></td>
              <td>Write data to TCP connection (alias for tcp_pathao)</td>
            </tr>
            <tr>
              <td><code>tcp_shuno</code></td>
              <td><code>connection</code></td>
              <td>Read data from TCP connection (async, returns promise)</td>
            </tr>
            <tr>
              <td><code>tcp_bondho</code></td>
              <td><code>connection</code></td>
              <td>Close TCP connection</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>UDP (User Datagram Protocol)</h2>

      <p>
        UDP is a connectionless protocol that&apos;s faster but doesn&apos;t guarantee delivery.
        Perfect for real-time applications like gaming, video streaming, and IoT where speed matters
        more than perfect reliability.
      </p>

      <h3>UDP Server</h3>

      <CodeBlock
        filename="udp_server.bang"
        code={`// UDP Server
dekho("Starting UDP server on port 9000...");

kaj handlePacket(packet) {
    dekho("Received from:", packet["remote_addr"]);
    dekho("Data:", packet["data"]);

    // Send response back to client
    udp_uttor(packet, "Received: " + packet["data"]);
}

udp_server_chalu(9000, handlePacket);
dekho("UDP server listening!");

jotokkhon (sotti) {
    ghumaao(1000);
}`}
      />

      <h3>UDP Client</h3>

      <CodeBlock
        filename="udp_client.bang"
        code={`proyash kaj udpClient() {
    dekho("Sending UDP packet...");

    // Send UDP packet (async)
    opekha udp_pathao("localhost", 9000, "Hello UDP!");

    dekho("Packet sent!");
}

udpClient();`}
      />

      <h3>UDP Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>udp_server_chalu</code></td>
              <td><code>port, handler</code></td>
              <td>Start UDP server with callback for each packet</td>
            </tr>
            <tr>
              <td><code>udp_pathao</code></td>
              <td><code>host, port, data</code></td>
              <td>Send UDP packet (async, returns promise)</td>
            </tr>
            <tr>
              <td><code>udp_uttor</code></td>
              <td><code>packet, data</code></td>
              <td>Send UDP response to client</td>
            </tr>
            <tr>
              <td><code>udp_shuno</code></td>
              <td><code>port, handler</code></td>
              <td>Listen for UDP packets (alias for udp_server_chalu)</td>
            </tr>
            <tr>
              <td><code>udp_bondho</code></td>
              <td><code>connection</code></td>
              <td>Close UDP connection</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>WebSocket</h2>

      <p>
        WebSocket provides full-duplex communication channels over a single TCP connection.
        Perfect for real-time applications like chat, live updates, and collaborative tools.
      </p>

      <h3>WebSocket Server</h3>

      <CodeBlock
        filename="websocket_chat.bang"
        code={`// WebSocket Chat Server
dekho("Starting WebSocket chat server on port 3000...");

bishwo clients = [];

kaj handleWebSocket(conn) {
    dekho("Client connected:", conn["remote_addr"]);

    // Add to clients list
    dhokao(clients, conn);

    // Broadcast messages to all clients
    dhoro message = conn["message"];
    ghuriye (dhoro i = 0; i < dorghyo(clients); i = i + 1) {
        websocket_pathao(clients[i], "Broadcast: " + message);
    }
}

websocket_server_chalu(3000, handleWebSocket);
dekho("WebSocket server running on ws://localhost:3000");

jotokkhon (sotti) {
    ghumaao(1000);
}`}
      />

      <h3>WebSocket Client</h3>

      <CodeBlock
        filename="websocket_client.bang"
        code={`proyash kaj wsClient() {
    dekho("Connecting to WebSocket...");

    // Connect to WebSocket server (async)
    dhoro ws = opekha websocket_jukto("ws://localhost:3000");
    dekho("Connected!");

    // Send messages
    websocket_pathao(ws, "Hello from BanglaCode!");
    websocket_pathao(ws, "This is message 2");

    // Wait a bit
    ghumaao(2000);

    // Close connection
    websocket_bondho(ws);
    dekho("Disconnected");
}

wsClient();`}
      />

      <h3>WebSocket Functions</h3>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Parameters</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><code>websocket_server_chalu</code></td>
              <td><code>port, handler</code></td>
              <td>Start WebSocket server with callback for messages</td>
            </tr>
            <tr>
              <td><code>websocket_jukto</code></td>
              <td><code>url</code></td>
              <td>Connect to WebSocket server (async, returns promise)</td>
            </tr>
            <tr>
              <td><code>websocket_pathao</code></td>
              <td><code>connection, message</code></td>
              <td>Send WebSocket message</td>
            </tr>
            <tr>
              <td><code>websocket_bondho</code></td>
              <td><code>connection</code></td>
              <td>Close WebSocket connection</td>
            </tr>
          </tbody>
        </table>
      </div>

      <h2>Connection Objects</h2>

      <p>
        All networking functions use connection objects (maps) to represent connections.
        These objects contain useful information about the connection:
      </p>

      <h3>TCP Connection Object</h3>

      <CodeBlock
        code={`{
    id: "tcp_conn_1",           // Unique connection ID
    host: "localhost",          // Host (client only)
    port: 8080,                 // Port (client only)
    remote_addr: "127.0.0.1:12345",  // Remote address
    local_addr: "127.0.0.1:8080",    // Local address
    data: "received data"       // Received data (server only)
}`}
      />

      <h3>UDP Packet Object</h3>

      <CodeBlock
        code={`{
    id: "udp_conn_1",           // Unique packet ID
    data: "packet data",        // Packet data
    remote_addr: "127.0.0.1:12345",  // Sender address
    local_addr: "127.0.0.1:9000"     // Receiver address
}`}
      />

      <h3>WebSocket Connection Object</h3>

      <CodeBlock
        code={`{
    id: "ws_conn_1",            // Unique connection ID
    url: "ws://localhost:3000", // WebSocket URL (client only)
    connected: sotti,           // Connection status
    remote_addr: "127.0.0.1:12345",  // Remote address
    local_addr: "127.0.0.1:3000",    // Local address
    message: "received message", // Received message (server only)
    type: "text"                // Message type (text/binary/close)
}`}
      />

      <h2>Function Name Meanings</h2>

      <div className="overflow-x-auto my-4">
        <table>
          <thead>
            <tr>
              <th>Function</th>
              <th>Bengali</th>
              <th>Meaning</th>
            </tr>
          </thead>
          <tbody>
            <tr><td><code>chalu</code></td><td>চালু</td><td>start/run</td></tr>
            <tr><td><code>jukto</code></td><td>যুক্ত</td><td>connect/join</td></tr>
            <tr><td><code>pathao</code></td><td>পাঠাও</td><td>send</td></tr>
            <tr><td><code>lekho</code></td><td>লেখো</td><td>write</td></tr>
            <tr><td><code>shuno</code></td><td>শোনো</td><td>listen/hear</td></tr>
            <tr><td><code>bondho</code></td><td>বন্ধ</td><td>close</td></tr>
            <tr><td><code>uttor</code></td><td>উত্তর</td><td>reply/response</td></tr>
          </tbody>
        </table>
      </div>

      <h2>Best Practices</h2>

      <ul>
        <li><strong>Use TCP</strong> when you need reliable, ordered delivery (HTTP APIs, file transfers, chat)</li>
        <li><strong>Use UDP</strong> when speed matters more than reliability (gaming, video streaming, IoT sensors)</li>
        <li><strong>Use WebSocket</strong> for real-time bidirectional communication (live chat, notifications, collaborative editing)</li>
        <li>Always use <code>proyash</code>/<code>opekha</code> for async operations</li>
        <li>Remember to close connections with <code>tcp_bondho</code>, <code>udp_bondho</code>, or <code>websocket_bondho</code></li>
        <li>Use error handling with <code>chesta</code>/<code>dhoro_bhul</code> for network operations</li>
      </ul>

      <h2>Complete Examples</h2>

      <p>
        Check out the <code>examples/</code> folder in the BanglaCode repository for complete working examples:
      </p>

      <ul>
        <li><code>tcp_server.bang</code> - TCP echo server</li>
        <li><code>tcp_client.bang</code> - TCP client with async/await</li>
        <li><code>udp_server.bang</code> - UDP server with responses</li>
        <li><code>udp_client.bang</code> - UDP client</li>
        <li><code>websocket_chat.bang</code> - WebSocket chat server</li>
        <li><code>websocket_client.bang</code> - WebSocket client</li>
      </ul>

      <h2>Next Steps</h2>

      <p>
        Learn more about related topics:
      </p>

      <ul>
        <li><Link href="/docs/async-await" className="text-primary hover:underline">Async/Await</Link> - Understanding promises and async operations</li>
        <li><Link href="/docs/http-server" className="text-primary hover:underline">HTTP Server</Link> - Building HTTP servers</li>
        <li><Link href="/docs/builtins" className="text-primary hover:underline">Built-in Functions</Link> - All 95+ built-in functions</li>
        <li><Link href="/docs/examples" className="text-primary hover:underline">Examples</Link> - More code examples</li>
      </ul>

      <DocNavigation currentPath="/docs/networking" />
    </div>
  );
}
