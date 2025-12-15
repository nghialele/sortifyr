import { useTrackGetHistory } from "@/lib/api/track";
import { formatDate } from "@/lib/utils";
import { Table } from "../molecules/Table";

export const TrackHistoryTable = () => {
  const { history, isLoading, isFetchingNextPage, hasNextPage, fetchNextPage } = useTrackGetHistory()

  const handleBottom = () => {
    if (!hasNextPage) return
    if (isFetchingNextPage) return

    fetchNextPage()
  }

  return (
    <Table
      idAccessor="historyId"
      columns={[
        { accessor: "name", title: "Track" },
        { accessor: "playedAt", render: ({ playedAt }) => <p className="text-muted">{formatDate(playedAt)}</p> }
      ]}
      records={history}
      fetching={isLoading || isFetchingNextPage}
      onScrollToBottom={handleBottom}
    />
  )
}
