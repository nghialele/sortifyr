import { PlaylistSchema } from "@/lib/types/playlist";
import { ActionIcon } from "@mantine/core";
import { LuX } from "react-icons/lu";
import { PlaylistCover } from "../playlist/PlaylistCover";

type Props = {
  playlist: PlaylistSchema;
  onDelete: (playlist: PlaylistSchema) => void;
}

export const DirectoryTreePlaylistEditable = ({ playlist, onDelete }: Props) => {
  return (
    <div className="flex items-center gap-2">
      <PlaylistCover playlist={playlist} />
      <p className="line-clamp-1 overflow-hidden text-ellipsis break-words">{playlist.name}</p>
      <p className="text-muted text-sm">{playlist.trackAmount}</p>
      <ActionIcon onClick={() => onDelete(playlist)} variant="transparent" className="text-muted ml-auto">
        <LuX />
      </ActionIcon>
    </div>
  )
}
