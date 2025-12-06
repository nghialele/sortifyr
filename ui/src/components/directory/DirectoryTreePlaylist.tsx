import { PlaylistSchema } from "@/lib/types/playlist";
import { PlaylistCover } from "../playlist/PlaylistCover";

type Props = {
  playlist: PlaylistSchema;
}

export const DirectoryTreePlaylist = ({ playlist }: Props) => {
  return (
    <div className="flex items-center gap-2">
      <PlaylistCover playlist={playlist} />
      <p className="line-clamp-1 overflow-hidden text-ellipsis break-words">{playlist.name}</p>
      <p className="text-muted text-sm">{playlist.trackAmount}</p>
    </div>
  )
}
