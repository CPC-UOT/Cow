<template>
  <div class="container">
    <VueMonacoEditor
      language="cpp"
      v-model:value="store.state.code"
      :theme="editorTheme"
      :options="MONACO_EDITOR_OPTIONS"
      @mount="handleMount"
      @keydown="formatCode()"
    />
  </div>
</template>
<script>
import { VueMonacoEditor } from "@guolao/vue-monaco-editor";
import { computed } from "vue";
import store from "../store/index.js";
export default {
  name: "monaco-view",
  components: { VueMonacoEditor },
  data() {
    return { store };
  },
};
</script>
<style scoped>
.container {
  margin-top: 2rem;
  height: 25rem;
  width: 70rem;
  overflow: hidden;
  padding-top: 1.5rem;
  background-color: #1e1e1e;
  border-radius: 2rem;
  border: #6835f8 5px solid;
}
</style>
<script setup>
import { shallowRef } from "vue";
const editorTheme = computed(() => "vs-dark");
const MONACO_EDITOR_OPTIONS = {
  automaticLayout: true,
  formatOnType: true,
  formatOnPaste: true,
  minimap: { enabled: false },
};
const editorRef = shallowRef();
const handleMount = (editor) => (editorRef.value = editor);
const theme = null;
console.log(theme);

// your action
console.log(shallowRef());
function formatCode() {
  editorRef.value?.getAction("editor.action.formatDocument").run();
}
</script>
