import { apiUrl } from "./axios";

// sse.js
export function initSSE(uid, callback) {
    const eventSource = new EventSource(apiUrl+'/events/' + uid);

    eventSource.onmessage = function(event) {
        callback(event.data);
    };

    eventSource.onerror = function() {
        console.error("SSE 连接错误");
    };
}
