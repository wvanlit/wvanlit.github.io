export interface JobData {
  title: string;
  period: string;
  description: string;
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
    description: `Building Studio from scratch, a document management product with AI-powered analysis baked in.`,
    bullets: [
      `I work across the whole stack, from document workflows and review tooling down to the infrastructure a brand-new product needs to get off the ground.`,
      `I keep our internal coding harnesses and dev workflows fast and pleasant to work in. They're the tools the rest of the team builds on.`,
      `I build the agent infrastructure in the app layer, the part that lets AI analysis and human review work together instead of fighting each other.`,
    ],
  },
  {
    title: "Senior Software Engineer @ Coolblue",
    period: "October 2024 - March 2026",
    description: `The engineer behind Coolblue's "ordered today, delivered tomorrow" promise, owning the customer-facing delivery journeys across our TypeScript and C# systems.`,
    bullets: [
      `Led a 6-month rewrite of the delivery proposition microservice from C# to TypeScript. Moving to pre-fetched caches and event-driven data took P95 from 250ms to 50ms and P99 from 500ms to 100ms, let it scale to millions of requests on a handful of ECS tasks, and put ownership back in the team's hands.`,
      `Designed an event-driven product data aggregator that gave the delivery apps one consistent read model instead of stitching several sources together, which cut both the integration work and the data mismatches that came with it.`,
    ],
  },
  {
    title: "Software Engineer @ Coolblue",
    period: "August 2022 - October 2024",
    description: `Built and looked after the shipment infrastructure behind 20,000+ shipments a day, modernizing legacy systems and shipping new C# microservices along the way.`,
    bullets: [
      `Designed and built a shipment event store on AWS (Lambda, SQS, SNS, DynamoDB) that became the source of truth for other domains and powered shipment status in the order overview.`,
      `Teamed up with a data analyst on a microservice that sets the insurance value of a shipment, running ONNX models written in Python directly inside .NET so we could make the call with ML in production.`,
    ],
  },
  {
    title: "Software Consultant @ ICT Group",
    period: "May 2021 - August 2022",
    description: `A consulting role delivering logistics software for clients in manufacturing and port operations.`,
    subProjects: [
      {
        title: "Project 2 - Nederlandse Spoorwegen (Dutch Railways)",
        bullets: [
          `Led the migration of a Manufacturing Execution System (MES) front-end off Internet Explorer and onto Chrome and ES6 JavaScript, which finally unblocked modern development on it.`,
          `Built workflow tools in TypeScript and React for factory-floor employees, cutting manual steps out of their day.`,
        ],
      },
      {
        title: "Project 1 - PSA International",
        bullets: [
          `Improved the developer tooling (Java, Spring Boot, Angular) used to validate systems after an update and catch regressions before they shipped.`,
          `Built an Azure CI/CD pipeline with Packer, Terraform, and Python that made spinning up and testing older system versions quick and repeatable.`,
        ],
      },
    ],
  },
  {
    title: "Graduate Intern @ ICT Group",
    period: "August 2020 - February 2021",
    description: `My graduation internship, focused on getting CI/CD working for embedded and IoT devices.`,
    bullets: [
      `Wrote my thesis on bringing CI/CD pipelines to embedded devices, and built a working prototype to show it could be done.`,
      `Got over-the-air updates running on ESP8266 microcontrollers through Azure CI/CD with C++ and Python, proving the approach held up in practice.`,
    ],
  },
  {
    title: "Research & Development Engineer @ Monta",
    period: "March 2020 - August 2020",
    description: `A hands-on R&D role using hardware and software prototypes to speed up warehouse picking.`,
    bullets: [
      `Researched whether light (lamps and LED strips) could make warehouse picking faster.`,
      `Wrote C# to drive a DMX spotlight that lit up the right products for pickers.`,
      `Built a C# UWP app on a Raspberry Pi that talked to Teensy microcontrollers over RabbitMQ.`,
      `Programmed the Teensy boards in C++ to point pickers at the right shelves with LED strips.`,
    ],
  },
  {
    title: "Technical Product Owner & Analyst @ Billink",
    period: "February 2019 - February 2020",
    description: `Started as a data analyst and grew into the technical product owner, leading two remote developers while keeping the credit check fast and reliable.`,
    bullets: [
      `Dug through data in Tableau and SQL to track down bugs and find where we could do better, both operationally and technically.`,
      `Built an asynchronous Go microservice with a custom rules engine to replace the old credit check, which made it a lot more reliable.`,
    ],
  },
  {
    title: "Peer Coach @ Hogeschool Rotterdam",
    period: "August 2018 - August 2019",
    description: `Helped CS teachers and first-year students with programming and math.`,
    bullets: [
      `Ran workshops and tutored students to make the intimidating topics feel more approachable.`,
      `Helped organize open days, hackathons, and welcome weeks to build community.`,
    ],
  },
];
