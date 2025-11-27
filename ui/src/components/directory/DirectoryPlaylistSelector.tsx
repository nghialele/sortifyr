import { Playlist } from "@/lib/types/playlist"
import { useDraggable } from "@dnd-kit/core";
import { CSS } from '@dnd-kit/utilities';

type Props = {
  playlists: Playlist[];
}

export const DirectoryPlaylistSelector = ({ playlists }: Props) => {
  return (
    <div className="border-2 border-gray-400 rounded-xl p-4 flex flex-wrap gap-2 h-fit">
      {playlists.map(p => <Entry key={p.id} playlist={p} />)}
    </div>
  )
}

const Entry = ({ playlist }: { playlist: Playlist }) => {
  const { attributes, listeners, setNodeRef, transform } = useDraggable({
    id: playlist.id,
  })

  const style = {
    transform: CSS.Translate.toString(transform),
  }

  return (
    <div ref={setNodeRef} className="border border-gray-100 z-10 shadow-xs bg-white rounded-md p-4 cursor-pointer" style={style} {...listeners} {...attributes}>
      <p className="font-bold">{playlist.name}</p>
      <p>{playlist.tracks}</p>
    </div>
  )
}

