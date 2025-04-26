<template>
  <div id="app">
    <TerminalResume />
  </div>
</template>

<script>
import TerminalResume from './components/TerminalResume.vue'

export default {
  name: 'App',
  components: {
    TerminalResume
  }
}
</script>

<style>
/* Reset all browser defaults */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body, html {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  overflow: hidden;
}

#app {
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}
</style>