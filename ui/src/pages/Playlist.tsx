import { LoadingSpinner } from "@/components/molecules/LoadingSpinner"
import { usePlaylistGetAll, usePlaylistSync } from "@/lib/api/playlist"
import { Playlist } from "@/lib/types/playlist"
import { Button } from "@mantine/core"
import { notifications } from "@mantine/notifications"

export const Playlists = () => {
  const { data: playlists, isLoading } = usePlaylistGetAll()

  const playlistSync = usePlaylistSync()

  if (isLoading) return <LoadingSpinner />

  const handleSync = () => {
    playlistSync.mutate(undefined, {
      onSuccess: () => notifications.show({ variant: "succes", message: "Syncing" }),
    })
  }

  return (
    <div>
      <Button onClick={handleSync}>
        Sync
      </Button>
      <div className="grid grid-cols-4">
        {playlists?.map(p => <Entry key={p.id} playlist={p} />)}
      </div>
    </div>
  )
}

const Entry = ({ playlist }: { playlist: Playlist }) => {
  return (
    <div className="border border-gray-200 shadow-md">
      <p className="font-bold">{playlist.name}</p>
      <p>{`Tracks: ${playlist.tracks}`}</p>
    </div>
  )
}
