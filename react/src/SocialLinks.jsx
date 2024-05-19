import { useEffect } from "react";
import { FaGithub, FaTwitter, FaLinkedin } from "react-icons/fa";

export default function SocialLinks({ user, className }) {
    const socialLinks = [
      { name: 'GitHub', url: user.GithubUrl, icon: <FaGithub/> },
      { name: 'LinkedIn', url: user.LinkedinUrl, icon: <FaLinkedin/> },
      { name: 'Twitter', url: user.TwitterUrl, icon: <FaTwitter/> },
    ];
    useEffect(()=> {
        console.log(user)
    },[]);
  
    return (
      <div className={className}>
        <div className="flex space-x-4">
          {socialLinks.map((link, index) => (
            link.url && (
              <a key={index} href={`https://${link.url}`} target="_blank" rel="noopener noreferrer">
                {link.icon}
              </a>
            )
          ))}
        </div>
      </div>
    );
  }
  