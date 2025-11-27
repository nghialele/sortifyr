import { LoadingSpinner } from "@/components/molecules/LoadingSpinner"
import { useDirectoryGetAll } from "@/lib/api/directory"
import { Directory } from "@/lib/types/directory"
import { Group, RenderTreeNodePayload, Tree, TreeNodeData } from "@mantine/core"
import { useMemo } from "react"
import { FaFolder, FaFolderOpen, FaPlay } from "react-icons/fa6"

export const Directories = () => {
  const { data: directories, isLoading } = useDirectoryGetAll()

  if (isLoading) return <LoadingSpinner />

  return (
    <div className="grid grid-cols-2">
      {directories?.map(d => <Entry key={d.id} directory={d} />)}
    </div>
  )
}

interface FileIconProps {
  isFolder: boolean;
  expanded: boolean;
}

const FileIcon = ({ isFolder, expanded }: FileIconProps) => {
  if (!isFolder) {
    return <FaPlay />
  }

  if (expanded) {
    return <FaFolderOpen />
  }

  return <FaFolder />
}

const Leaf = ({ node, expanded, hasChildren, elementProps }: RenderTreeNodePayload) => {
  return (
    <Group gap={5} {...elementProps}>
      <FileIcon isFolder={hasChildren} expanded={expanded} />
      <span>{node.label}</span>
    </Group>
  )
}

const toData = (directory: Directory): TreeNodeData => {
  return {
    label: directory.name,
    value: directory.name,
    children: directory.children?.map(toData) ?? [],
  }
}

const Entry = ({ directory }: { directory: Directory }) => {
  const data: TreeNodeData[] = useMemo(() => {
    return [toData(directory)];
  }, [directory])

  return (
    <Tree
      selectOnClick
      clearSelectionOnOutsideClick
      data={data}
      renderNode={(payload) => <Leaf {...payload} />}
    />
  )
}
