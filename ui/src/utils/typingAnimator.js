// src/utils/typingAnimator.js

export class TypingAnimator {
  constructor(options = {}) {
    this.typingSpeed = {
      command: options.commandSpeed || 100,
      content: options.contentSpeed || 5
    };
    this.pauseDuration = {
      afterCommand: options.afterCommandPause || 500,
      afterContent: options.afterContentPause || 0
    };
  }

  /**
   * Types out a command character by character
   * @param {string} command - The command to type
   * @param {function} onCharacter - Callback for each character typed
   * @param {function} onComplete - Callback when typing is complete
   * @param {number} initialDelay - Delay before starting to type
   */
  typeCommand(command, onCharacter, onComplete, initialDelay = 500) {
    let i = 0;
    let typedText = '';
    
    const typeNextChar = () => {
      if (i < command.length) {
        typedText += command.charAt(i);
        onCharacter(typedText);
        i++;
        setTimeout(typeNextChar, this.typingSpeed.command);
      } else {
        setTimeout(() => {
          onComplete();
        }, this.pauseDuration.afterCommand);
      }
    };
    
    setTimeout(typeNextChar, initialDelay);
  }

  /**
   * Types out content in chunks for faster display
   * @param {string} content - The content to type
   * @param {function} onChunk - Callback for each chunk typed
   * @param {function} onComplete - Callback when typing is complete
   * @param {number} chunkSize - Characters per chunk
   */
  typeContent(content, onChunk, onComplete, chunkSize = 10) {
    let visibleLength = 0;
    const contentLength = content.length;
    
    const typeNextChunk = () => {
      if (visibleLength < contentLength) {
        visibleLength = Math.min(visibleLength + chunkSize, contentLength);
        const visibleContent = content.substring(0, visibleLength);
        onChunk(visibleContent);
        setTimeout(typeNextChunk, this.typingSpeed.content);
      } else {
        setTimeout(() => {
          onComplete();
        }, this.pauseDuration.afterContent);
      }
    };
    
    typeNextChunk();
  }
}