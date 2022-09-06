const SERVER_URL = "http://127.0.0.1:8080/counter";

export function fetchCount() {
  return new Promise<{ data: number }>((resolve) =>
    fetch(SERVER_URL)
    .then((response) => response.json())
    .then((json) => resolve(json))
  );
}
