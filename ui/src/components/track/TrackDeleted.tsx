import { useTrackGetDeleted } from "@/lib/api/track"
import { TrackFilter } from "@/lib/types/track"
import { formatDate } from "@/lib/utils"
import { Table } from "../molecules/Table"
import { SectionTitle } from "../atoms/Page"
import { usePlaylistGetAll } from "@/lib/api/playlist"
import { useState } from "react"
import { Select } from "../molecules/Select"

export const TrackDeleted = () => {
  const { data: playlists, isLoading: isLoadingPlaylists } = usePlaylistGetAll()

  const [filter, setFilter] = useState<TrackFilter>({})

  const { tracks, isLoading, isFetchingNextPage, hasNextPage, fetchNextPage } = useTrackGetDeleted(filter)

  const handleBottom = () => {
    if (!hasNextPage) return
    if (isFetchingNextPage) return

    fetchNextPage()
  }

  return (
    <>
      <SectionTitle
        title="Recently Deleted"
        description="An overview of recently deleted tracks from playlists."
      />
      <Select
        data={playlists?.map(p => ({ value: p.id.toString(), label: p.name }))}
        value={filter.playlistId?.toString()}
        onChange={(v) => setFilter({ ...filter, playlistId: v ? v : undefined })}
        placeholder="Filter track by playlist..."
        disabled={isLoadingPlaylists}
      />
      <Table
        columns={[
          { accessor: "name", title: "Track", width: "45%", ellipsis: true },
          { accessor: "playlist.name", title: "Playlist", width: "35%", ellipsis: true },
          { accessor: "deletedAt", render: ({ deletedAt }) => <p className="text-muted">{formatDate(deletedAt)}</p> }
        ]}
        records={tracks}
        noRecordsText="No recorded deleted tracks yet"
        fetching={isLoading || isFetchingNextPage}
        onScrollToBottom={handleBottom}
      />
    </>
  )
}
