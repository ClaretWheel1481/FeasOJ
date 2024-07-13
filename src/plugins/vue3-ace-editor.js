import { VAceEditor } from 'vue3-ace-editor';
import 'ace-builds/src-noconflict/mode-json'; // 加载语言定义文件
import 'ace-builds/src-noconflict/theme-chrome'; // 加载主题定义文件
import modeJava from "ace-builds/src-noconflict/mode-java?url";
import modeCpp from "ace-builds/src-noconflict/mode-c_cpp?url";
import modeGolang from "ace-builds/src-noconflict/mode-golang?url";
import modePy from "ace-builds/src-noconflict/mode-python?url";
import { config } from "ace-builds";

config.setModuleUrl("ace/mode/java", modeJava);
config.setModuleUrl("ace/mode/c_cpp", modeCpp);
config.setModuleUrl("ace/mode/golang", modeGolang);
config.setModuleUrl("ace/mode/python", modePy);

export default {
  components: {
    VAceEditor,
  },
  data() {
    return {
      content: JSON.stringify({ message: 'Hello Ace' }),
    };
  },
};
