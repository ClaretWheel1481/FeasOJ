import { apiUrl } from "./axios";
import { i18n } from '../plugins/vue-i18n';

// sse
export function initSSE(uid, callback) {
    function connect() {
        const eventSource = new EventSource(apiUrl + '/notification/' + uid);

        eventSource.onmessage = function (event) {
            callback(event.data);
        };

        eventSource.onerror = function () {
            console.error(i18n.global.t("message.sseError"));
            eventSource.close();
            setTimeout(connect, 3000);  // 3秒后重连
        };
    }
    
    connect();
}
