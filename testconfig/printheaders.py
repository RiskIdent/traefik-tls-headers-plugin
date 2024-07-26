import http.server
import socketserver


class MyRequestHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        response_body = "\n".join([f"{header}: {value}" for header, value in self.headers.items()])

        self.send_response(200)
        self.end_headers()
        self.wfile.write(response_body.encode("utf-8"))


PORT = 8888
if __name__ == "__main__":
    with socketserver.TCPServer(("", PORT), MyRequestHandler) as httpd:
        print(f"Serving on port {PORT}")
        httpd.serve_forever()
