export interface JobData {
  title: string;
  period: string;
  description?: string;
  descriptionHtml?: string;
  bullets?: string[];
  subProjects?: {
    title: string;
    bullets: string[];
  }[];
}

export const jobs: JobData[] = [
  {
    title: "Senior Software Engineer @ RobotX",
    period: "March 2026 - Present",
    descriptionHtml: `Working on <a href="https://robotx.com/products/studio/fsx" target="_blank" rel="noopener noreferrer">Studio</a>, a document management product with AI analysis built in.`,
    bullets: [
      `I work across the whole stack, from document workflows and review tooling to the infrastructure a new product needs to go from 0 to 1.`,
      `I keep our internal coding harnesses and dev workflows fast and usable, because the rest of the team builds on top of them.`,
      `I build the agent layer in the app, which lets AI analysis and human review fit into the same workflow without getting in each other's way.`,
    ],
  },
  {
    title: "Senior Software Engineer @ Coolblue",
    period: "October 2024 - March 2026",
    description: `Owned the customer-facing delivery journeys behind Coolblue's "ordered today, delivered tomorrow" promise across TypeScript and C# systems.`,
    bullets: [
      `Led a six-month rewrite of the delivery proposition microservice from C# to TypeScript. Pre-fetched caches and event-driven data brought P95 down from 250ms to 50ms and P99 from 500ms to 100ms, while scaling to millions of requests on a handful of ECS tasks.`,
      `Designed an event-driven product data aggregator that gave the delivery apps one read model instead of piecing several sources together, which cut integration work and reduced data mismatches.`,
    ],
  },
  {
    title: "Software Engineer @ Coolblue",
    period: "August 2022 - October 2024",
    description: `Built and maintained the shipment infrastructure behind 20,000+ shipments a day, while modernizing legacy systems and shipping new C# microservices.`,
    bullets: [
      `Designed and built a shipment event store on AWS (Lambda, SQS, SNS, DynamoDB) that became the source of truth for other domains and powered shipment status in the order overview.`,
      `Worked with a data analyst on a microservice that sets shipment insurance value, running ONNX models written in Python directly inside .NET so we could use ML in production.`,
    ],
  },
  {
    title: "Software Consultant @ ICT Group",
    period: "May 2021 - August 2022",
    description: `Built logistics software for clients in manufacturing and port operations.`,
    subProjects: [
      {
        title: "Project 2 - Nederlandse Spoorwegen (Dutch Railways)",
        bullets: [
          `Led the migration of a Manufacturing Execution System (MES) front end off Internet Explorer and onto Chrome and ES6 JavaScript, which finally made modern development on it possible.`,
          `Built workflow tools in TypeScript and React for factory-floor employees, cutting out manual steps.`,
        ],
      },
      {
        title: "Project 1 - PSA International",
        bullets: [
          `Improved the developer tooling (Java, Spring Boot, Angular) used to validate systems after updates and catch regressions before release.`,
          `Built an Azure CI/CD pipeline with Packer, Terraform, and Python that made spinning up and testing older system versions quick and repeatable.`,
        ],
      },
    ],
  },
  {
    title: "Graduate Intern @ ICT Group",
    period: "August 2020 - February 2021",
    description: `Graduation internship focused on getting CI/CD working for embedded and IoT devices.`,
    bullets: [
      `Wrote my thesis on bringing CI/CD pipelines to embedded devices and built a working prototype to show the approach was viable.`,
      `Got over-the-air updates running on ESP8266 microcontrollers through Azure CI/CD with C++ and Python.`,
    ],
  },
  {
    title: "Research & Development Engineer @ Monta",
    period: "March 2020 - August 2020",
    description: `Built hardware and software prototypes to make warehouse picking faster.`,
    bullets: [
      `Researched whether light, using lamps and LED strips, could make warehouse picking faster.`,
      `Wrote C# to drive a DMX spotlight that lit up the right products for pickers.`,
      `Built a C# UWP app on a Raspberry Pi that talked to Teensy microcontrollers over RabbitMQ.`,
      `Programmed the Teensy boards in C++ to point pickers at the right shelves with LED strips.`,
    ],
  },
  {
    title: "Technical Product Owner & Analyst @ Billink",
    period: "February 2019 - February 2020",
    description: `Started as a data analyst and grew into the technical product owner, leading two remote developers and looking after the credit check flow.`,
    bullets: [
      `Dug through data in Tableau and SQL to track down bugs and find where the process could be improved, both operationally and technically.`,
      `Built an asynchronous Go microservice with a custom rules engine to replace the old credit check and make it more reliable.`,
    ],
  },
  {
    title: "Peer Coach @ Hogeschool Rotterdam",
    period: "August 2018 - August 2019",
    description: `Helped CS teachers and first-year students with programming and math.`,
    bullets: [
      `Ran workshops and tutored students to make the intimidating topics less intimidating.`,
      `Helped organize open days, hackathons, and welcome weeks to build community.`,
    ],
  },
];
