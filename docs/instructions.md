# Project overview
We are building a system to keep track of an organizations application portfiolio. Mid-size & enterprise organizations use many applications. Application Portfolio Management (APM) helps manage these, providing insights into the number, value, technologies, and the lifecycle of applications.

Most competitors are still very labor intensive, not very user-friendly and outdated. We want to change this

We need to build a proof of concept for this application. 
- It can import software from a CSV file.
- It has a pluggable classification engine which will be fleshed out later. We will use data found on the internet to enrich the imported data.
- It has a UI to manage these things where different user types can log in. The user role we start with is the application portfolio manager.

The tech stack to be used:
- A backend in golang
- Postgres DB
- SQLC for talking to the DB.
- react for the frontend
- Tailwind CSS and DaisyUI for styling the frontend and interactive elements

Libs used
https://phoenixnap.com/kb/postgresql-kubernetes
for k8s we want to use Kustomize and flux

## Software Types & Subtypes

### Introduction

Software is a collection of programs, data, and instructions that tell a computer how to perform specific tasks or operations. It includes applications, operating systems, and utilities that help users interact with computer hardware and perform various functions.

### Software types

There are two types of software that can be distinguished:

- Application Software is designed for end-users to perform specific tasks or activities. Examples include word processors (e.g., Microsoft Word), web browsers (e.g., Google Chrome), games, and business applications (e.g., CRM systems). Application software is built on top of system software and utilizes the resources and services provided by the operating system to perform its functions.
- System Software is designed to manage and control the hardware components of a computer and provide a platform for running application software. It includes operating systems (e.g., Windows, macOS, Linux), device drivers, and utilities that help in system management and maintenance. System software acts as an intermediary between the hardware and user applications, ensuring smooth and efficient operation of the entire computer system.

2.1 Application Software Types

Application Software Types

Descriptions

Application Software Types

Descriptions

Productivity Software

Productivity software includes tools designed to help users create documents, manage data, and present information. These applications streamline tasks and improve efficiency in both professional and personal settings. Examples include:

Word Processors: Software like Microsoft Word or Google Docs for creating and editing text documents.

Spreadsheets: Tools like Microsoft Excel or Google Sheets for managing numerical data and performing calculations.

Presentation Software: Applications like Microsoft PowerPoint or Google Slides for creating visual presentations.

Creative Software

Creative Software: Creative software encompasses tools used for designing graphics, editing videos, composing music, and other artistic endeavors. These applications are essential for professionals in creative industries as well as hobbyists. Examples include:

Graphic Design: Adobe Photoshop or Illustrator for creating and editing images and illustrations.

Video Editing: Adobe Premiere Pro or Final Cut Pro for editing and producing videos.

Music Production: Ableton Live or Logic Pro for composing and producing music.

Business Software

Business Software: Business software includes applications that help organizations manage various aspects of their operations, such as customer relationships, financial transactions, and resource planning. Examples include:

Customer Relationship Management (CRM): Salesforce or HubSpot for managing customer interactions and data.

Enterprise Resource Planning (ERP): SAP or Oracle ERP for integrating and managing core business processes.

Accounting Software: QuickBooks or Xero for tracking financial transactions and managing accounts.

Educational Software

Educational Software: Educational software is designed to facilitate learning and instruction. These applications are used in schools, universities, and for self-study. Examples include:

Language Learning: Duolingo or Rosetta Stone for learning new languages.

Online Courses: Coursera or Khan Academy for accessing a wide range of educational courses and resources.

Classroom Management: Google Classroom or Moodle for organizing and managing classroom activities and assignments.

Communication Software

Communication Software: Communication software includes tools that enable users to exchange information and collaborate. These applications are essential for personal and professional communication. Examples include:

Email Clients: Microsoft Outlook or Gmail for managing and sending emails.

Messaging Apps: Slack or Microsoft Teams for real-time messaging and collaboration.

Video Conferencing: Zoom or Skype for conducting virtual meetings and video calls.

Entertainment Software

Entertainment Software: Entertainment software provides users with leisure and recreational activities. These applications include games, streaming services, and multimedia players. Examples include:

Streaming Services: Netflix or Spotify for streaming movies, TV shows, and music.

Video Games: Steam or Xbox Game Pass for accessing and playing video games.

Multimedia Players: VLC Media Player or Windows Media Player for playing audio and video files.

Utility Software

Utility Software: Utility software includes tools that help manage and optimize computer systems. These applications perform tasks such as system maintenance, security, and data compression. Examples include:

Antivirus Software: Norton or McAfee for protecting computers from malware and viruses.

File Compression: WinRAR or 7-Zip for compressing and decompressing files.

Backup Software: Acronis or Backblaze for creating and managing data backups.

Web Browsers

Web Browsers: Web browsers are applications used for accessing and navigating the internet. These tools allow users to view web pages, download files, and interact with online content. Examples include:

Google Chrome: A widely used web browser known for its speed and extensions.

Mozilla Firefox: An open-source browser with a focus on privacy and customization.

Microsoft Edge: A browser integrated with Windows, offering performance and security features.


2.2 System Software Types

System Software Types

Descriptions

Integrations

Integrations are software tools and solutions that enable different applications, systems, or services to work together seamlessly. They facilitate the exchange of data and functionality between disparate software components, allowing them to operate as a cohesive system. Integrations can be achieved through various means, such as APIs (Application Programming Interfaces), middleware, and connectors. The goal of integrations is to improve efficiency, streamline workflows, and enhance the overall functionality of the combined systems.

Development Frameworks

Development Frameworks: These are comprehensive platforms that provide a structured environment for building software applications. Frameworks include predefined classes, functions, and tools that help developers manage common tasks and enforce best practices. They often dictate the architecture of an application and provide a foundation for code organization, reducing the complexity of development. Frameworks also offer various utilities for different stages of development, such as testing, debugging, and deployment. Examples of development frameworks include Angular (for building web applications), Django (for web development in Python), and Ruby on Rails (for web development in Ruby).

Development Libraries

Development Libraries: These are collections of pre-written code that developers can use to perform common tasks, functions, or operations within their applications. Libraries help streamline the development process by providing reusable code components, which can save time and effort. They are designed to be flexible and modular, allowing developers to choose specific functions without enforcing a particular structure or architecture on the entire application. Examples of development libraries include React.js (for building user interfaces), lodash (for utility functions), and jQuery (for simplifying HTML DOM manipulation).

Drivers

Drivers are specialized software that allow the operating system and other software to communicate with hardware devices. They act as intermediaries, translating the high-level commands from the operating system into low-level instructions that the hardware can understand. Drivers are essential for the proper functioning of hardware components, such as printers, graphics cards, network adapters, and more. Without the appropriate drivers, the hardware devices may not work correctly or may not be recognized by the operating system.

Operating Systems

Operating Systems (OS) are essential system software that manage computer hardware and software resources, and provide common services for computer programs. An OS acts as an intermediary between users and the computer hardware, enabling the execution of application software. Key functions of an operating system include process management, memory management, file system management, and device management. Popular examples of operating systems include Windows, macOS, Linux, and Android.

Firmware

Firmware is a type of software that is embedded in hardware devices to control and manage their operations. It provides low-level control for the device's specific hardware and is typically stored in non-volatile memory, such as ROM, EEPROM, or flash memory. Firmware is essential for the basic functioning of various devices, such as computers, smartphones, routers, and embedded systems, and it often includes a set of instructions or code necessary for the hardware to perform its intended functions.

Storage systems

Storage systems refer to software solutions that facilitate the efficient management, organization, retrieval, and storage of data. This includes software designed for databases, caching mechanisms, and data storage systems. Storage software ensures that data is stored securely, efficiently, and is readily accessible when needed, optimizing overall data management and improving performance across various applications and systems.

Embedded Systems

Embedded Systems are specialized computer systems that are designed to perform dedicated functions within larger systems or devices. Embedded systems are typically integrated into hardware and operate with real-time constraints. They consist of a combination of hardware and software, where the software (firmware) is embedded into the hardware components. Common examples of embedded systems include microcontrollers in appliances, control systems in automobiles, medical devices, and IoT (Internet of Things) devices. These systems are optimized for specific tasks, ensuring efficiency, reliability, and real-time performance.

## Categories of software


# Core functionalities
# Object Functionality and Description

## Users

### Login
- A user should be able to log in.

### Register
- A user should be able to register/create an account.

### MFA
- A user should be able to set up two-factor authentication for their account.

### User Management
A user should be able to manage the users in their organization:
- Add new users to the organization.
- Remove users from the organization.
- Change roles of the users in the organization.

#### Roles for the MVP:
- **Organization Admin**: Able to add/remove users.
- **Application Portfolio Manager**: Should not be able to access user management (but can see the rest).
- **Stakeholder**: Should not be able to log in/access anything.

### Delete Account
- A user should be able to delete their own account.
- If the user is the **Organization Admin**, then:
  - They can only delete their own account if there is another **Organization Admin**.
  - If no other accounts exist, deleting their account also deletes the entire organization.

### Manage Profile
A user should be able to edit:
- First name
- Last name
- Avatar/Picture
- Email address

## Organization

### Manage Organization
- Adjust the display name of their organization.
- When a new account is created, the user must specify the organization (workspace) name.
- Creating an account results in the creation of a new account and organization.

### Customized Subdomain
- When a new account/organization is created, it should link to a separate subdomain (e.g., `https://organisationname.unicorn.pm`).

## Applications

### Create
- A user should be able to create a new application for their organization.

### Bulk Import
- A user should be able to import multiple applications via a CSV file (specific format) for their organization.

### Overview
- A user should be able to view and search through all uploaded applications within their organization.

### Detail
- A user should be able to access the detail page of a specific application to view its specifics.

### AI Enrichment
- The system should automatically enrich application data for COTS applications when missing:
  - Description
  - Applicable categories (functionalities)
  - Applicable Type/Subtype

### Master Application Database
- All COTS applications added to an organization should create a shadow copy in a master record (not visible to the end user).
- All organizational COTS applications should be related to a master application record.
- The master record contains only public information.
- Organizations can override the default master record details (e.g., name, description).

### Clustering
- A user should be able to create an **application cluster** for their organization.
- They should be able to:
  - Specify details of the cluster.
  - Select applications to include in the cluster.
- An application can belong to multiple clusters (or none).

## Entities

### Create
- A user should be able to create an **Entity** (e.g., organization or company) for their organization.
- The entity name should not be changeable, except when aligning with the related master entity record.

### Bulk Import
- A user should be able to bulk import **Entities** via CSV (specific format) for their organization.

### Overview
- A user should be able to view all listed entities in their organization.

### Detail
- A user should be able to view details of their organization's entities.

### Master Entity Database
- For each created entity, a **master entity record** should be created and related to the organization’s entity.

## News

### Scrape News
- For each master record (applications and related entities), news articles should be scraped daily from public sources.

## Ranking

### Scrape Ranking
- For each application master record, the ranking should be scraped daily (or weekly) from public ranking sources.

## Competitors

### Scrape Competitors
- For each application master record, competitor solutions should be scraped automatically.
- This should trigger additional master records to be created (if they don’t already exist).
- Relationships should be created between competitor records.

## Reports

### Duplication Report
- A user should be able to access a **Duplication Report**.
- The report should visualize applications with:
  - Identical functionalities (overlap)
  - Direct competitors in the organization's portfolio

## Logs

### Access Logs
- Every user login should create an access log entry, linked to the user and their role at that time.

### Edit Logs
- Every change a user makes to any entity should generate a log entry including:
  - User who made the change
  - Affected entity
  - Old value
  - New value


# Doc

# Current file structure