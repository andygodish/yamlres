// src/services/resumeService.js
import axios from 'axios';

export class ResumeService {
  /**
   * Fetches resume data from the API
   * @returns {Promise<Object>} Resume data
   */
  static async fetchResumeData() {
    try {
      console.log('Attempting to fetch resume data...');
      const response = await axios.get('/api/resume');
      console.log('Resume data received:', response.data);
      return response.data;
    } catch (error) {
      console.error('Failed to fetch resume data:', error);
      throw new Error(`Failed to fetch resume: ${error.message}`);
    }
  }

  /**
   * Extracts username from resume basics
   * @param {Object} resume - Resume data
   * @returns {string} Username
   */
  static extractUsername(resume) {
    if (!resume?.basics?.name) return 'user';
    
    const nameParts = resume.basics.name.toLowerCase().split(' ');
    if (nameParts.length > 0) {
      return nameParts[0].replace(/[^a-z0-9]/g, '');
    }
    return 'user';
  }

  /**
   * Validates resume data structure
   * @param {Object} resume - Resume data to validate
   * @returns {boolean} True if valid
   */
  static validateResumeData(resume) {
    if (!resume || typeof resume !== 'object') {
      return false;
    }
    
    // Check for required sections (at least basics should exist)
    return Boolean(resume.basics);
  }
}