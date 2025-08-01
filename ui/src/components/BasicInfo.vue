<template>
  <div class="basic-info">
    <div class="line">{{ basics.name }}</div>
    <div class="line" v-if="basics.label">{{ basics.label }}</div>
    <div class="line" v-if="basics.email">{{ basics.email }}</div>
    <div class="line" v-if="basics.phone">{{ basics.phone }}</div>
    <div class="line" v-if="basics.url">
      <a :href="basics.url" target="_blank" rel="noopener noreferrer">{{ basics.url }}</a>
    </div>
    <div class="line" v-if="basics.summary">{{ basics.summary }}</div>
    <div class="line" v-if="locationString">{{ locationString }}</div>
    <div class="profiles" v-if="basics.profiles && basics.profiles.length">
      <div v-for="profile in basics.profiles" :key="profile.network" class="profile">
        <a :href="profile.url" target="_blank" rel="noopener noreferrer">
          {{ profile.network }}: {{ profile.username }}
        </a>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BasicInfo',
  props: {
    basics: {
      type: Object,
      required: true
    }
  },
  mounted() {
  console.log('BasicInfo received:', JSON.stringify(this.basics, null, 2));
  },
  computed: {
    locationString() {
      if (!this.basics.location) return '';
      const loc = this.basics.location;
      const parts = [];
      
      if (loc.city) parts.push(loc.city);
      if (loc.region) parts.push(loc.region);
      if (loc.countryCode) parts.push(loc.countryCode);
      
      return parts.join(', ');
    }
  }
};
</script>

<style scoped>
.basic-info {
  margin: 15px 0 25px;
  padding-left: 10px;
  border-left: 3px solid #E95420;  /* Ubuntu orange */
}

.basic-info .line {
  margin-bottom: 5px;
}

.basic-info a {
  color: #729FCF;  /* Ubuntu terminal blue */
  text-decoration: none;
}

.basic-info a:hover {
  text-decoration: underline;
}

.profiles {
  margin-top: 10px;
}

.profile {
  margin-bottom: 3px;
}
</style>