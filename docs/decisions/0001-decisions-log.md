# Decisions log

## Context

This is a reference project with the practices and approaches to Software
Development. One of the practices is to track project decisions in a traceable
way. In a real project, the lack of such a process can lead to potential
misunderstandings, longer onboarding time, and challenges in maintaining a
common view of the system.

## Decision

The project decisions will be logged in this repo. This log will be called
**Decisions Logs** and will be usedd as a practice for documenting significant
project decisions.

### Decisions log rules

- Decisions will be kept in the project repository under
  `docs/decisions/NNNN-title-in-kebab-case.md`, and will be numbered
  sequentially and monotonically.
- Decisions will keep a collection of the project's significant decisions: that
  affect technology vision, architecture, and engineering practices.
- This log will follow a more relaxed set of rules than **ADRs**.
  - Existing decisions can be updated with more details and/or clarifications.
  - New decisions can overwrite old ones but must link to the previous to make
    the log traceable.
  - Related decisions should be linked to make a decision tree.
- A decision document should follow the next structure.
  - **Title**: A document has a title that is a short noun phrase. For example,
    "Public endpoints should adhere to REST" or "Web frontend should be built
    with React and Next.js".
  - **Context**: This section describes the forces at play, including
    technological, organizational, and team's local. The language in this
    section is value-neutral. It is simply describing facts.
  - **Decision**: This section describes the response to these forces.
  - **Consequences**: This section describes the resulting context. Positive and
    negative consequences should be listed here.

### Inspiration and examples

- [Architectural Descision Records](https://adr.github.io/)
- [Documenting Architecture Decisions](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)
- [[Github] Architecture decision record examples](https://github.com/joelparkerhenderson/architecture-decision-record)
- [[AWS Doc] ADR process](https://docs.aws.amazon.com/prescriptive-guidance/latest/architectural-decision-records/adr-process.html)

## Consequences

### Positive

- **Transparency**: Decisions log will provide a record of the rationale behind
  decisions, making it easier for team members to understand the thoughts behind
  choices.
- **Knowledge Base**: Decisions log will serve for onboarding new team members
  and providing a knowledge base, ensuring that everyone is on the same page
  regarding the project.
- **Communication**: Decisions log will support collaboration among team members
  by providing a centralized repository for technical and process decisions that
  can be easily referenced. Also, Decisions log helps to develop writing vs.
  verbal communication.

### Challenges

- **Team Engagement**: If there is no clear understanding of the value of this
  practice, the team members may ignore this documentation.
- **Lack of Value**: Without a clear and documented process for creating and
  reviewing decision records, this documentation may become unmaintainable.
- **Documentation Overhead**: The process of creating and maintaining the
  Decisions Log can introduce additional overhead, especially with the
  fast-changing requirements and environment.
- **Outdated Records**: Outdated or incomplete Decisions Log can mislead team
  members and create confusion about the current state of the system.
