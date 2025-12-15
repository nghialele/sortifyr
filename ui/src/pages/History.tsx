import { Page, PageTitle, Section } from "@/components/atoms/Page";
import { TrackHistoryTable } from "@/components/track/TrackHistoryTable";

export const History = () => {
  return (
    <Page>
      <PageTitle
        title="Recently Played"
        description="An overview of recently played tracks"
      />

      <Section>
        <TrackHistoryTable />
      </Section>
    </Page>
  )
}
