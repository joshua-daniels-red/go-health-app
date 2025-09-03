import requests
import sys
import json
import os

BASE_URL = os.getenv("GO_SERVER_URL", "http://localhost:8080")

def fetch_all_movies(page_size=2):
    """
    Generator that yields movies from /data endpoint in paginated chunks.
    """
    page = 1
    while True:
        resp = requests.get(f"{BASE_URL}/data", params={"page": page, "page_size": page_size})
        if resp.status_code != 200:
            break
        items = resp.json()
        if not items:
            break
        yield from items
        page += 1


def transform_imdb():
    results = []
    for movie in fetch_all_movies():
        old = movie["imdb"]
        new = round(old + 0.1, 1)
        results.append({
            "id": movie["id"],
            "change": f"imdb score will be changed from {old} to {new}"
        })
    return results


def transform_title():
    results = []
    for movie in fetch_all_movies():
        old = movie["movie_title"]
        new = f"{old} (English)"
        results.append({
            "id": movie["id"],
            "change": f"title will be changed from '{old}' to '{new}'"
        })
    return results


def main():
    if len(sys.argv) < 2:
        print(json.dumps({"error": "no operation provided"}))
        sys.exit(1)

    op = sys.argv[1]
    if op == "imdb":
        output = transform_imdb()
    elif op == "title":
        output = transform_title()
    else:
        output = {"error": f"unsupported operation: {op}"}

    print(json.dumps(output))


if __name__ == "__main__":
    main()
