import json
from http.server import BaseHTTPRequestHandler, HTTPServer


class RouterHandler(BaseHTTPRequestHandler):
    routes = {
        "/callback": "callback",
    }

    def do_GET(self):
        path = self.path.split("?")[0]  # Usuń query string

        if path in self.routes:
            handler = getattr(self, f"handle_{self.routes[path]}")
            handler()
        else:
            self.send_error(404, "Not Found")

    def handle_callback(self):
        self.send_response(200)
        self.send_header("Content-type", "application/json")
        self.end_headers()
        data = {"message": "Hello from API", "status": "ok"}
        self.wfile.write(json.dumps(data).encode())


if __name__ == "__main__":
    server = HTTPServer(("localhost", 8080), RouterHandler)
    print("Serwer uruchomiony na http://localhost:8080")
    server.serve_forever()
