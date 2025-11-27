import { DirectorySchema } from "@/lib/types/directory"
import { DirectoryNode } from "./DirectoryNode"

type Props = {
  root: DirectorySchema;
  onUpdate: (directory: DirectorySchema) => void;
  onDelete: (directory: DirectorySchema) => void;
}

export const DirectoryTree = ({ root, onUpdate, onDelete }: Props) => {
  return (
    <DirectoryNode
      directory={root}
      onUpdate={onUpdate}
      onDelete={onDelete}
      level={0}
    />
  )
}

