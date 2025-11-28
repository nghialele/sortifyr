import { Playlist } from "@/lib/types/playlist"
import { ComponentProps } from "react";
import { FaRegCirclePlay } from "react-icons/fa6";
import { LoadableImage } from "../atoms/LoadableImage";
import { cn } from "@/lib/utils";

type Props = {
  playlist: Pick<Playlist, "id" | "hasCover">;
} & Pick<ComponentProps<"div">, "className">

export const PlaylistCover = ({ playlist, className }: Props) => {
  if (playlist.hasCover) {
    return <LoadableImage src={`/api/playlist/cover/${playlist.id}`} className={cn("w-8 h-8 rounded-md", className)} />
  }

  return <FaRegCirclePlay className={cn("text-green-500 w-8 h-8", className)} />
}
