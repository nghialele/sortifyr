import { useDirectoryGetAll } from "@/lib/api/directory";
import { Directory } from "@/lib/types/directory";
import { Playlist } from "@/lib/types/playlist";
import { Group, RenderTreeNodePayload, Tree, TreeNodeData } from "@mantine/core";
import { useMemo } from "react";
import { FaRegCirclePlay, FaRegFolder, FaRegFolderOpen } from "react-icons/fa6";
import { LoadingSpinner } from "../molecules/LoadingSpinner";
import { PlaylistCover } from "./PlaylistCover";

interface FileIconProps {
  isFolder: boolean;
  playlist?: Playlist;
  expanded: boolean;
}

const FileIcon = ({ playlist, isFolder, expanded }: FileIconProps) => {
  if (!isFolder) {
    if (playlist) return <PlaylistCover playlist={playlist} className="w-6 h-6" />
    else return <FaRegCirclePlay className="w-6" />
  }

  if (expanded) {
    return <FaRegFolderOpen />
  }

  return <FaRegFolder />
}

const Leaf = ({ node, expanded, hasChildren, elementProps }: RenderTreeNodePayload) => {
  return (
    <Group py={2} {...elementProps}>
      <div className="flex items-center gap-2">
        <FileIcon playlist={node.nodeProps?.playlist} isFolder={hasChildren} expanded={expanded} />
        <span className="whitespace-nowrap">{node.label}</span>
      </div>
    </Group>
  )
}

const toData = (directory: Directory): TreeNodeData => {
  const playlists = directory.playlists.map(p => ({ label: p.name, value: p.name, nodeProps: { "playlist": p } }))

  return {
    label: directory.name,
    value: directory.name,
    children: [...playlists, ...(directory.children?.map(toData) ?? [])],
  }
}

export const PlaylistTreeView = () => {
  const { data: directories, isLoading } = useDirectoryGetAll()

  const data = useMemo(() => directories?.map(toData) ?? [], [directories])
  console.log(data)

  if (isLoading) return <LoadingSpinner />

  return (
    <Tree
      selectOnClick
      clearSelectionOnOutsideClick
      data={data}
      renderNode={payload => <Leaf {...payload} />}
    />
  )
}
