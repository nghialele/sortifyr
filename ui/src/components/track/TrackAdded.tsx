import { useTrackGetAdded } from "@/lib/api/track"
import { TrackFilter } from "@/lib/types/track"
import { formatDate } from "@/lib/utils"
import { Table } from "../molecules/Table"
import { SectionTitle } from "../atoms/Page"
import { useState } from "react"
import { usePlaylistGetAll } from "@/lib/api/playlist"
import { Select } from "../molecules/Select"

export const TrackAdded = () => {
  const { data: playlists, isLoading: isLoadingPlaylists } = usePlaylistGetAll()

  const [filter, setFilter] = useState<TrackFilter>({})

  const { tracks, isLoading, isFetchingNextPage, hasNextPage, fetchNextPage } = useTrackGetAdded(filter)

  const handleBottom = () => {
    if (!hasNextPage) return
    if (isFetchingNextPage) return

    fetchNextPage()
  }

  return (
    <>
      <SectionTitle
        title="Recently Added"
        description="An overview of recently added tracks to playlists."
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
          { accessor: "createdAt", render: ({ createdAt }) => <p className="text-muted">{formatDate(createdAt)}</p> }
        ]}
        records={tracks}
        noRecordsText="No recored tracks added yet"
        fetching={isLoading || isFetchingNextPage || isLoadingPlaylists}
        onScrollToBottom={handleBottom}
      />
    </>
  )
}
