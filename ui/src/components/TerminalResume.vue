<template>
  <div class="terminal-fullscreen">
    <a 
      :href="githubTagUrl" 
      class="version-badge" 
      target="_blank" 
      rel="noopener noreferrer" 
      title="View this version on GitHub"
      @click="openGithubTag"
    >
      v{{ version }}
    </a>
    <div class="terminal-content" ref="terminalContent">
      <template v-if="stage === 'typing-command'">
        <div class="command-line">
          <span class="prompt">yamlres:~$</span> 
          <span class="typing-command">{{ typedCommand }}</span>
          <span class="cursor" :class="{ 'blink': typingCommandComplete && !loadingData }">_</span>
        </div>
      </template>
      
      <div v-else-if="stage === 'error'" class="error-line">
        <span class="prompt">yamlres:~$</span> 
        <span class="command">cat resume.yaml</span>
        <div class="error-message">{{ error }}</div>
      </div>
      
      <template v-else-if="stage === 'typing-content' || stage === 'complete'">
        <!-- Command shown at top -->
        <div class="command-line">
          <span class="prompt">yamlres:~$</span> 
          <span class="command">cat resume.yaml</span>
        </div>
        
        <!-- Typed content that appears gradually -->
        <div ref="typedContent" v-html="typedContent"></div>
        
        <!-- Cursor that appears at the end of typed content when not complete -->
        <span v-if="stage === 'typing-content'" class="cursor blink">_</span>
      </template>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'TerminalResume',
  data() {
    return {
      version: process.env.VUE_APP_VERSION || 'X.Y.Z',
      repoUrl: 'https://github.com/andygodish/yamlres',
      username: 'user',
      resume: null,
      stage: 'typing-command', // stages: typing-command, loading-data, typing-content, complete, error
      error: null,
      commandToType: 'cat resume.yaml',
      typedCommand: '',
      typedContent: '',
      fullContent: '',
      typingCommandComplete: false,
      loadingData: false,
      typingSpeed: {
        command: 100, // ms per character for command
        content: 5    // ms per character for content (much faster)
      },
      pauseDuration: {
        afterCommand: 500, // ms to pause after command typed
        afterContent: 0    // ms to pause after content typed
      }
    };
  },
  computed: {
        githubTagUrl() {
        return `${this.repoUrl}/tree/v${this.version}`;
      }
  },
  created() {
    this.startTypingCommand();
  },
  methods: {
    startTypingCommand() {
      // Start the command typing animation
      let i = 0;
      const typeNextChar = () => {
        if (i < this.commandToType.length) {
          this.typedCommand += this.commandToType.charAt(i);
          i++;
          setTimeout(typeNextChar, this.typingSpeed.command);
        } else {
          // Command typing is complete, pause briefly then load data
          this.typingCommandComplete = true;
          setTimeout(() => {
            this.loadingData = true;
            this.fetchResumeData();
          }, this.pauseDuration.afterCommand);
        }
      };
      
      setTimeout(typeNextChar, 500); // Initial delay before typing starts
    },
    
    async fetchResumeData() {
      try {
        console.log('Attempting to fetch resume data...');
        const response = await axios.get('/api/resume');
        console.log('Resume data received:', response.data);
        
        this.resume = response.data;
        
        // Set username from resume if available
        if (this.resume && this.resume.basics && this.resume.basics.name) {
          const nameParts = this.resume.basics.name.toLowerCase().split(' ');
          if (nameParts.length > 0) {
            this.username = nameParts[0].replace(/[^a-z0-9]/g, '');
          }

          // Emit event with the full name to update document title
          this.$emit('resume-loaded', this.resume.basics.name);
        }
        
        // Generate HTML content for the resume
        this.generateResumeContent();
        
        // Start typing out the content
        this.stage = 'typing-content';
        this.startTypingContent();
        
      } catch (err) {
        console.error('Failed to fetch resume data:', err);
        this.error = `Error: ${err.message}`;
        this.stage = 'error';
      }
    },
    
    generateResumeContent() {
      // This function generates the HTML content for the resume
      let content = '';
      
      // Basic info
      if (this.resume.basics) {
        content += `<div class="basic-info">
          <div class="line">${this.resume.basics.name}</div>
          <div class="line">${this.resume.basics.email}</div>
        </div>`;
      }
      content += `<div class="section-divider"></div>`;

      // Skills
      if (this.resume.skills && this.resume.skills.length > 0) {
        content += `<div class="section-header">TECH STACK</div>`;
        content += `<div class="section">`;
        
        this.resume.skills.forEach(skillGroup => {
          content += `<div class="skill-group">
            <div class="skill-group-name">${skillGroup.name}:</div>
            <div class="skill-items">`;
            
          skillGroup.keywords.forEach(skill => {
            content += `<span class="skill-item">${skill}</span>`;
          });
          
          content += `</div>
          </div>`;
        });
        
        content += `</div>`;
      }
      
      content += `<div class="section-divider"></div>`;

      // Work experience
      if (this.resume.work && this.resume.work.length > 0) {
        content += `<div class="section-header">WORK EXPERIENCE</div>`;
        content += `<div class="section">`;
        
        this.resume.work.forEach(job => {
          content += `<div class="job-entry">
            <div class="job-title">${job.company} (${job.startDate} - ${job.endDate})</div>
            <div class="job-position">${job.position}</div>
            <ul class="job-highlights">`;
            
          job.highlights.forEach(highlight => {
            content += `<li>${highlight}</li>`;
          });
          
          content += `</ul>
          </div>`;
        });
        
        content += `</div>`;
      }
      
      // Education
      content += `<div class="section-divider"></div>`;
      if (this.resume.education && this.resume.education.length > 0) {
        content += `<div class="section-header">EDUCATION</div>`;
        content += `<div class="section">`;
          
        this.resume.education.forEach(edu => {
          content += `<div class="education-entry">
            <div class="education-title">${edu.institution} (${edu.startDate} - ${edu.endDate})</div>
            <div class="education-type">${edu.studyType}</div>
            </div>`;
          });
          
        content += `</div>`;
      }
      content += `<div class="section-divider"></div>`;

      // Certificates
      if (this.resume.certificates && this.resume.certificates.length > 0) {
        content += `<div class="section-header">CERTIFICATIONS</div>`;
        content += `<div class="section">
          <ul class="certificates-list">`;
          
        this.resume.certificates.forEach(cert => {
          content += `<li class="certificate-item">${cert.name}</li>`;
        });
        
        content += `</ul>
        </div>`;
      }
      content += `<div class="section-divider"></div>`;
             
      // Volunteer
      if (this.resume.volunteer && this.resume.volunteer.length > 0) {
        content += `<div class="section-header">VOLUNTEER</div>`;
        content += `<div class="section">`;
        
        this.resume.volunteer.forEach(vol => {
          content += `<div class="volunteer-entry">
            <div class="volunteer-title">${vol.organization} (${vol.startDate} - ${vol.endDate})</div>
          </div>`;
        });
        
        content += `</div>`;
      }
      
      content += `<div class="section-divider"></div>`;
      
      this.fullContent = content;
    },
    
    startTypingContent() {
      let contentLength = this.fullContent.length;
      let visibleLength = 0;
      const chunkSize = 10; // Characters to add per tick (for faster typing)
      
      const typeNextChunk = () => {
        if (visibleLength < contentLength) {
          visibleLength = Math.min(visibleLength + chunkSize, contentLength);
          this.typedContent = this.fullContent.substring(0, visibleLength);
          
          // Scroll to keep up with typing
          this.$nextTick(() => {
            if (this.$refs.terminalContent) {
              this.$refs.terminalContent.scrollTop = this.$refs.terminalContent.scrollHeight;
            }
          });
          
          setTimeout(typeNextChunk, this.typingSpeed.content);
        } else {
          // Content typing is complete
          setTimeout(() => {
            this.stage = 'complete';
          }, this.pauseDuration.afterContent);
        }
      };
      
      // Start typing content
      typeNextChunk();
    }
  }
};
</script>

<style>
/* Ubuntu-inspired color scheme with radial gradient */
.terminal-fullscreen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle at center, #4C1D3F 0%, #3C1434 30%, #300A24 60%, #1E0517 100%);
  color: #FFFFFF;
  font-family: 'Ubuntu Mono', monospace;
  z-index: 1000;
  overflow: auto;
}

.terminal-content {
  width: 90%;
  max-width: 900px;
  margin: 60px auto 100px; /* Added bottom margin */
  line-height: 1.6;
  padding: 0 20px;  /* Added horizontal padding */
}

.prompt {
  color: #4E9A06;  /* Ubuntu terminal green */
  font-weight: bold;
  margin-right: 8px;
}

.command, .typing-command {
  color: #729FCF;  /* Ubuntu terminal blue */
}

.command-line {
  margin-bottom: 45px;
  min-height: 24px; /* Ensures consistent height during typing */
}

.cursor {
  display: inline-block;
  color: #FFFFFF;
  margin-left: 1px;
  font-weight: bold;
}

.cursor.blink {
  animation: blink 1s step-end infinite;
}

.error-message {
  color: #FF5555;
  margin-top: 5px;
}

.section-header {
  color: #E95420;  /* Ubuntu orange */
  font-size: 1.2em;
  font-weight: bold;
  margin: 25px 0 15px;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.section {
  margin: 20px 0;
}

.section-divider {
  border-top: 1px dashed #5E2750;  /* Ubuntu terminal border */
  margin: 30px 0;
}

.basic-info {
  margin: 15px 0 25px;
  padding-left: 10px;
  border-left: 3px solid #E95420;  /* Ubuntu orange */
}

.basic-info .line {
  margin-bottom: 5px;
}

.job-entry, .education-entry, .volunteer-entry {
  margin-bottom: 20px;
  /* padding-left: 10px; */
}

.job-title, .education-title, .volunteer-title {
  color: #E95420;  /* Ubuntu orange */
  font-weight: bold;
}

.job-position, .education-type {
  color: #AEA79F;  /* Ubuntu light gray */
  margin: 5px 0;
}

.job-highlights {
  padding-left: 13px;
  margin: 8px 0;
}

.job-highlights li {
  margin-bottom: 5px;
}

.skill-group {
  margin-bottom: 20px;
}

.skill-group-name {
  color: #E95420;  /* Ubuntu orange */
  font-weight: bold;
  margin-bottom: 10px;
}

.skill-items {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.skill-item {
  background-color: rgba(255, 255, 255, 0.1);
  padding: 5px 10px;
  border-radius: 4px;
  font-size: 14px;
}

.certificates-list {
  list-style-type: none;
}

.certificate-item {
  margin-bottom: 5px;
  position: relative;
}

.certificate-item::before {
  margin-right: 8px;
}

.end-message {
  padding: 10px 0 30px;
  color: #FFFFFF;
  font-style: italic;
}

.version-badge {
  position: fixed;
  top: 15px;
  right: 15px;
  background-color: rgba(233, 84, 32, 0.7); /* Semi-transparent Ubuntu orange */
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
  background-color: rgba(233, 84, 32, 0.9); /* Slightly more opaque on hover */
  transform: scale(1.05);
}

.version-badge:active {
  transform: scale(0.98);
}

@keyframes blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}

/* Add Ubuntu Mono font */
@import url('https://fonts.googleapis.com/css2?family=Ubuntu+Mono&display=swap');

/* Add responsive styling */
@media (max-width: 768px) {
  .terminal-content {
    width: 95%;
    margin: 30px auto 60px;
    padding: 0 15px;
  }
  
  .skill-items {
    gap: 5px;
  }
}
</style>