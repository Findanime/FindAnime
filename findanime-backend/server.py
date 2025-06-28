from fastapi import FastAPI, Request
import uvicorn
from chatgpt.recommender import AskAI
from starlette.middleware.base import BaseHTTPMiddleware
from fastapi.middleware.cors import CORSMiddleware

class LoggingMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        print(f"Incoming request: {request.method} {request.url}")
        response = await call_next(request)
        print(f"Response status: {response.status_code}")
        return response

class AnimeRecommenderService:
    def __init__(self):
        pass

    def suggest_anime(self, anime: str) -> dict:
        try:
            response = AskAI(anime)
            return {"success": True, "error": False, "msg": response}
        except Exception as e:
            return {"success": False, "error": True, "msg": str(e)}

class AnimeRecommenderAPI:
    def __init__(self):
        self.app = FastAPI()
        self.service = AnimeRecommenderService()
        self._add_routes()
        self._add_middlewares()

    def _add_routes(self):
        self.app.get("/api/suggest")(self.suggest)

    async def suggest(self, anime: str):
        return self.service.suggest_anime(anime)

    def _add_middlewares(self):
        self.app.add_middleware(LoggingMiddleware)
        self.app.add_middleware(
            CORSMiddleware,
            allow_origins=["*"],  # Accept requests from any origin
            allow_credentials=True,
            allow_methods=["*"],  # Allow all HTTP methods
            allow_headers=["*"],  # Allow all headers
        )

    def run(self, host="0.0.0.0", port=80):
        uvicorn.run(self.app, host=host, port=port)

if __name__ == "__main__":
    api = AnimeRecommenderAPI()
    api.run()