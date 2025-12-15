import { useTrackGetDeleted } from "@/lib/api/track"
import { TrackFilter } from "@/lib/types/track"
import { formatDate } from "@/lib/utils"
import { Table } from "../molecules/Table"

type Props = {
  filter?: TrackFilter
}

export const TrackDeletedTable = ({ filter }: Props) => {
  const { tracks, isLoading, isFetchingNextPage, hasNextPage, fetchNextPage } = useTrackGetDeleted(filter)

  const handleBottom = () => {
    if (!hasNextPage) return
    if (isFetchingNextPage) return

    fetchNextPage()
  }

  return (
    <Table
      columns={[
        { accessor: "name", title: "Track", width: "45%", ellipsis: true },
        { accessor: "playlist.name", title: "Playlist", width: "35%", ellipsis: true },
        { accessor: "deletedAt", render: ({ deletedAt }) => <p className="text-muted">{formatDate(deletedAt)}</p> }
      ]}
      records={tracks}
      fetching={isLoading || isFetchingNextPage}
      onScrollToBottom={handleBottom}
    />
  )
}
