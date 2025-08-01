// Updated BasicInfo.vue with typing animation

<template>
  <div class="basic-info">
    <!-- Fields appear progressively based on typingProgress -->
    <div class="line" v-if="shouldShow('name')">{{ basics.name }}</div>
    <div class="line" v-if="shouldShow('label') && basics.label">{{ basics.label }}</div>
    <div class="line" v-if="shouldShow('email') && basics.email">{{ basics.email }}</div>
    <div class="line" v-if="shouldShow('phone') && basics.phone">{{ basics.phone }}</div>
    <div class="line" v-if="shouldShow('url') && basics.url">
      <a :href="basics.url" target="_blank" rel="noopener noreferrer">{{ basics.url }}</a>
    </div>
    <div class="line" v-if="shouldShow('summary') && basics.summary">{{ basics.summary }}</div>
    <div class="line" v-if="shouldShow('location') && locationString">{{ locationString }}</div>
    <div class="profiles" v-if="shouldShow('profiles') && basics.profiles && basics.profiles.length">
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
    },
    typingProgress: {
      type: Number,
      default: 0 // 0 = not started, 100 = complete
    }
  },
  data() {
    return {
      // Define the order and timing of field reveals
      fieldOrder: [
        { name: 'name', threshold: 5 },
        { name: 'label', threshold: 15 },
        { name: 'email', threshold: 25 },
        { name: 'phone', threshold: 35 },
        { name: 'url', threshold: 45 },
        { name: 'summary', threshold: 60 },
        { name: 'location', threshold: 80 },
        { name: 'profiles', threshold: 90 }
      ]
    };
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
  },
  methods: {
    shouldShow(fieldName) {
      const field = this.fieldOrder.find(f => f.name === fieldName);
      if (!field) return true; // Show by default if not in fieldOrder
      return this.typingProgress >= field.threshold;
    }
  },
  mounted() {
    console.log('BasicInfo received:', JSON.stringify(this.basics, null, 2));
    console.log('Initial typing progress:', this.typingProgress);
  },
  watch: {
    typingProgress(newProgress) {
      console.log('BasicInfo typing progress:', newProgress);
    }
  }
};
</script>

<style scoped>
.basic-info {
  margin: 15px 0 25px;
  padding-left: 10px;
  border-left: 3px solid #E95420;
}

.basic-info .line {
  margin-bottom: 5px;
}

.basic-info a {
  color: #729FCF;
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