export class ResumeContentGenerator {
  static generateResumeContent(resume) {
    let content = '';
    
    // Basic info
    if (resume.basics) {
      content += `<div class="basic-info">
        <div class="line">${resume.basics.name}</div>
        <div class="line">${resume.basics.email}</div>
      </div>`;
    }
    content += `<div class="section-divider"></div>`;

    // Skills
    content += this.generateSkillsSection(resume.skills);
    content += `<div class="section-divider"></div>`;

    // Work experience
    content += this.generateWorkSection(resume.work);
    
    // Education
    content += `<div class="section-divider"></div>`;
    content += this.generateEducationSection(resume.education);
    content += `<div class="section-divider"></div>`;

    // Certificates
    content += this.generateCertificatesSection(resume.certificates);
    content += `<div class="section-divider"></div>`;
             
    // Volunteer
    content += this.generateVolunteerSection(resume.volunteer);
    content += `<div class="section-divider"></div>`;
    
    return content;
  }

  static generateSkillsSection(skills) {
    if (!skills || skills.length === 0) return '';
    
    let content = `<div class="section-header">TECH STACK</div>`;
    content += `<div class="section">`;
    
    skills.forEach(skillGroup => {
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
    return content;
  }

  static generateWorkSection(work) {
    if (!work || work.length === 0) return '';
    
    let content = `<div class="section-header">WORK EXPERIENCE</div>`;
    content += `<div class="section">`;
    
    work.forEach(job => {
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
    return content;
  }

  static generateEducationSection(education) {
    if (!education || education.length === 0) return '';
    
    let content = `<div class="section-header">EDUCATION</div>`;
    content += `<div class="section">`;
      
    education.forEach(edu => {
      content += `<div class="education-entry">
        <div class="education-title">${edu.institution} (${edu.startDate} - ${edu.endDate})</div>
        <div class="education-type">${edu.studyType}</div>
        </div>`;
      });
      
    content += `</div>`;
    return content;
  }

  static generateCertificatesSection(certificates) {
    if (!certificates || certificates.length === 0) return '';
    
    let content = `<div class="section-header">CERTIFICATIONS</div>`;
    content += `<div class="section">
      <ul class="certificates-list">`;
      
    certificates.forEach(cert => {
      content += `<li class="certificate-item">${cert.name}</li>`;
    });
    
    content += `</ul>
    </div>`;
    return content;
  }

  static generateVolunteerSection(volunteer) {
    if (!volunteer || volunteer.length === 0) return '';
    
    let content = `<div class="section-header">VOLUNTEER</div>`;
    content += `<div class="section">`;
    
    volunteer.forEach(vol => {
      content += `<div class="volunteer-entry">
        <div class="volunteer-title">${vol.organization} (${vol.startDate} - ${vol.endDate})</div>
      </div>`;
    });
    
    content += `</div>`;
    return content;
  }
}