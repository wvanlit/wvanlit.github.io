import type { Company } from "../types/work";

export const companies: Company[] = [
  {
    name: "Coolblue",
    url: "https://www.coolblue.nl/",
    jobs: [
      {
        title: "Senior Software Engineer",
        period: { start: "2022-08", end: undefined }, // promoted Oct 2024
        projects: [
          {
            title: "Microservices platform ownership",
            bullets: [
              "Responsible for 20+ C# microservices enabling 20k parcels/day",
            ],
          },
          {
            title: "Eventstore for live track & trace",
            bullets: [
              "AWS Lambda, SQS, SNS, DynamoDB",
              ">1M events/day",
            ],
          },
          {
            title: "Fraud detection for shipping insurance",
            bullets: [
              "Built ML-powered microservice to automate insurance decisions",
            ],
          },
          {
            title: "Backoffice UIs for carriers",
            bullets: ["TypeScript, React"],
          },
        ],
      },
    ],
  },
  {
    name: "ICT Group",
    url: "https://ict.eu/",
    jobs: [
      {
        title: "Software Consultant",
        period: { start: "2021-05", end: "2022-08" },
        projects: [
          {
            title: "Nederlandse Spoorwegen (NS) — MES frontend migration",
            bullets: [
              "Migrated MES frontend from Internet Explorer to Chrome",
              "ES6 JavaScript, later TypeScript + React for workflow tools",
            ],
          },
          {
            title: "PSA International — Ops tools & CI/CD",
            bullets: [
              "Improved operational tools in Java (Spring Boot) & Angular",
              "Set up Azure CI/CD with Packer, Terraform, Python for test matrix",
            ],
          },
        ],
      },
      {
        title: "Graduation Intern",
        period: { start: "2020-08", end: "2021-02" },
        projects: [
          {
            title: "CI/CD for embedded devices",
            bullets: [
              "Thesis + working prototype",
              "OTA firmware updates on ESP8266 with Azure, C++, Python",
            ],
          },
        ],
      },
    ],
  },
  {
    name: "Monta",
    url: "https://monta.nl/",
    jobs: [
      {
        title: "Research & Development Engineer",
        period: { start: "2020-03", end: "2020-08" },
        projects: [
          {
            title: "Light-guided order picking",
            bullets: [
              "C# (cloud) and C (embedded) software for warehouse picking",
            ],
          },
        ],
      },
    ],
  },
  {
    name: "Billink",
    url: "https://billink.nl/",
    jobs: [
      {
        title: "Technical Product Owner & Analyst",
        period: { start: "2019-02", end: "2020-02" },
        projects: [
          {
            title: "Credit check microservice",
            bullets: ["Built in Go"],
          },
          {
            title: "Checkout & analytics",
            bullets: ["Tableau & Python for data analysis", "Led 2 external devs"],
          },
        ],
      },
    ],
  },
  {
    name: "Hogeschool Rotterdam",
    url: "https://www.hogeschoolrotterdam.nl/",
    jobs: [
      {
        title: "Peer Coach",
        period: { start: "2018-08", end: "2019-08" },
        projects: [
          {
            title: "Student coaching",
            bullets: [
              "Coached first-year CS students in programming & math",
              "Organized workshops, hackathons, campus events",
            ],
          },
        ],
      },
    ],
  },
];
