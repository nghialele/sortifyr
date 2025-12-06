import { DirectorySchema } from "@/lib/types/directory";
import { Stack, StackProps } from "@mantine/core";
import { LoadingSpinner } from "../molecules/LoadingSpinner";
import { DirectoryTreeNode } from "./DirectoryTreeNode";
import { DirectoryTreeNodeEditable } from "./DirectoryTreeNodeEditable";
import { cn } from "@/lib/utils";

type Props = {
  roots: DirectorySchema[];
  isLoading: boolean;
  title?: string;
  editable?: boolean;
  onUpdate?: (directory: DirectorySchema) => void;
  onDelete?: (directory: DirectorySchema) => void;
} & StackProps

export const DirectoryTree = ({ roots, isLoading, title, editable = false, onUpdate, onDelete, className, ...props }: Props) => {
  const child = () => {
    if (isLoading) return <LoadingSpinner />
    if (roots.length === 0) return null

    return (
      <>
        {roots.map(r => {
          if (!editable) return <DirectoryTreeNode key={r.iid} directory={r} level={0} />

          if (!onUpdate) onUpdate = () => null
          if (!onDelete) onDelete = () => null

          return <DirectoryTreeNodeEditable key={r.iid} directory={r} level={0} onUpdate={onUpdate} onDelete={onDelete} />
        })}
      </>
    )
  }

  return (
    <Stack p="md" gap="md" justify="start" bg="background.0" className={cn("overflow-auto rounded-xl select-none", className)} {...props}>
      {title && <p className="text-muted">{title}</p>}
      {child()}
    </Stack>
  )
}

