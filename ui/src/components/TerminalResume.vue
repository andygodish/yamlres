<template>
  <div class="terminal-fullscreen">
    <VersionBadge :version="version" :repo-url="repoUrl" />
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

        <div v-if="resume">
          <!-- Pass typing progress to BasicInfo -->
          <BasicInfo 
            v-if="resume.basics" 
            :basics="resume.basics" 
            :typing-progress="basicsTypingProgress" 
          />
          <div class="section-divider"></div>
    
          <div v-html="typedContent"></div>
        </div>
        
        <!-- Cursor that appears at the end of typed content when not complete -->
        <span v-if="stage === 'typing-content'" class="cursor blink">_</span>
      </template>
    </div>
  </div>
</template>

<script>
import { ResumeService } from '@/services/resumeService';
import { ResumeContentGenerator } from '@/utils/resumeContentGenerator';
import { TypingAnimator } from '@/utils/typingAnimator';
import VersionBadge from './VersionBadge.vue';
import BasicInfo from './BasicInfo.vue';

export default {
  name: 'TerminalResume',
  components: {
    BasicInfo,
    VersionBadge
  },
  data() {
    return {
      version: process.env.VUE_APP_VERSION || 'X.Y.Z',
      repoUrl: 'https://github.com/andygodish/yamlres',
      username: 'user',
      resume: null,
      stage: 'typing-command',
      error: null,
      commandToType: 'cat resume.yaml',
      typedCommand: '',
      typedContent: '',
      fullContent: '',
      typingCommandComplete: false,
      loadingData: false,
      basicsTypingProgress: 0,  // Added for BasicInfo typing animation
      typingAnimator: new TypingAnimator({
        commandSpeed: 100,
        contentSpeed: 5,
        afterCommandPause: 500,
        afterContentPause: 0
      })
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
      this.typingAnimator.typeCommand(
        this.commandToType,
        (typedText) => {
          this.typedCommand = typedText;
        },
        () => {
          this.typingCommandComplete = true;
          this.loadingData = true;
          this.fetchResumeData();
        }
      );
    },
    
    async fetchResumeData() {
      try {
        this.resume = await ResumeService.fetchResumeData();
        
        if (this.resume?.basics?.name) {
          this.username = ResumeService.extractUsername(this.resume);
          this.$emit('resume-loaded', this.resume.basics.name);
        }
        
        this.generateResumeContent();
        this.stage = 'typing-content';
        this.startTypingContent();
        
      } catch (err) {
        this.error = err.message;
        this.stage = 'error';
      }
    },
    
    generateResumeContent() {
      this.fullContent = ResumeContentGenerator.generateResumeContent(this.resume);
    },
    
    startTypingContent() {
      let contentLength = this.fullContent.length;
      let visibleLength = 0;
      const chunkSize = 10;
      
      const typeNextChunk = () => {
        if (visibleLength < contentLength) {
          visibleLength = Math.min(visibleLength + chunkSize, contentLength);
          this.typedContent = this.fullContent.substring(0, visibleLength);
          
          // Calculate progress for BasicInfo (0-100)
          this.basicsTypingProgress = Math.min((visibleLength / contentLength) * 100, 100);
          
          this.scrollToBottom();
          setTimeout(typeNextChunk, 5); // Using direct speed value instead of this.typingSpeed.content
        } else {
          this.basicsTypingProgress = 100; // Ensure it's complete
          setTimeout(() => {
            this.stage = 'complete';
          }, 0); // Using direct value instead of this.pauseDuration.afterContent
        }
      };
      
      typeNextChunk();
    },

    scrollToBottom() {
      this.$nextTick(() => {
        if (this.$refs.terminalContent) {
          this.$refs.terminalContent.scrollTop = this.$refs.terminalContent.scrollHeight;
        }
      });
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