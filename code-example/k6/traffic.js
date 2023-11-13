import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  const failure = Math.random() >= 0.99 ? true : false;
  const sleepMs = Math.random() >= 0.99 ? 100 : 0;

  http.get(`http://host.docker.internal:8081/simple-service/ping?fail=${failure}&sleep=${sleepMs}ms`);
  sleep(1);
}
