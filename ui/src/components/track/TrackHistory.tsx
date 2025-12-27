import { useTrackGetHistory } from "@/lib/api/track";
import { formatDate } from "@/lib/utils";
import { Table } from "../molecules/Table";
import { SectionTitle } from "../atoms/Page";

export const TrackHistory = () => {
  const { history, isLoading, isFetchingNextPage, hasNextPage, fetchNextPage } = useTrackGetHistory()

  const handleBottom = () => {
    if (!hasNextPage) return
    if (isFetchingNextPage) return

    fetchNextPage()
  }

  return (
    <>
      <SectionTitle
        title="Recently Played"
        description={`An overview of recently played tracks.\nRun the tracks task if a track has no title.`}
      />
      <Table
        idAccessor="historyId"
        columns={[
          { accessor: "name", title: "Track", width: "80%", ellipsis: true },
          { accessor: "playedAt", render: ({ playedAt }) => <p className="text-muted">{formatDate(playedAt)}</p> }
        ]}
        records={history}
        noRecordsText="No recorded tracks yet"
        fetching={isLoading || isFetchingNextPage}
        onScrollToBottom={handleBottom}
      />
    </>
  )
}
