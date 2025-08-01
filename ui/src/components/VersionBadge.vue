// src/components/VersionBadge.vue
<template>
  <a 
    :href="githubTagUrl" 
    class="version-badge" 
    target="_blank" 
    rel="noopener noreferrer" 
    :title="`View v${version} on GitHub`"
  >
    v{{ version }}
  </a>
</template>

<script>
export default {
  name: 'VersionBadge',
  props: {
    version: {
      type: String,
      required: true
    },
    repoUrl: {
      type: String,
      required: true
    }
  },
  computed: {
    githubTagUrl() {
      return `${this.repoUrl}/tree/v${this.version}`;
    }
  }
};
</script>

<style scoped>
.version-badge {
  position: fixed;
  top: 15px;
  right: 15px;
  background-color: rgba(233, 84, 32, 0.7);
  color: white;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.8em;
  font-weight: bold;
  z-index: 1001;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
  text-decoration: none;
  transition: background-color 0.2s, transform 0.1s;
}

.version-badge:hover {
  background-color: rgba(233, 84, 32, 0.9);
  transform: scale(1.05);
}

.version-badge:active {
  transform: scale(0.98);
}
</style>