import { apiUrl } from "./axios";
import { i18n } from '../plugins/vue-i18n';

// sse.js
export function initSSE(uid, callback) {
    const eventSource = new EventSource(apiUrl + '/events/' + uid);

    eventSource.onmessage = function (event) {
        callback(event.data);
    };

    eventSource.onerror = function () {
        console.error(i18n.global.t("message.sseError"));
    };
}
