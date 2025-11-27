import { useDirectoryGetAll, useDirectorySync } from "@/lib/api/directory";
import { convertDirectorySchema, DirectorySchema } from "@/lib/types/directory";
import { getUuid } from "@/lib/utils";
import { Button } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { FaPlus } from "react-icons/fa6";
import { Confirm } from "../molecules/Confirm";
import { notifications } from "@mantine/notifications";

type Props = {
  roots: DirectorySchema[];
  setRoots: (roots: DirectorySchema[]) => void;
}

export const DirectoryToolbar = ({ roots, setRoots }: Props) => {
  const { data: directories } = useDirectoryGetAll()

  const [openedReset, { open: openReset, close: closeReset }] = useDisclosure()
  const [openedSave, { open: openSave, close: closeSave }] = useDisclosure()

  const save = useDirectorySync()

  const handleDirectoryCreate = () => {
    const newRoot: DirectorySchema = {
      iid: getUuid(),
      name: "New Directory",
      playlists: [],
      children: []
    }

    const updated = [...roots, newRoot]
    setRoots(updated)
  }

  const handleResetInit = () => {
    openReset()
  }

  const handleReset = () => {
    setRoots(convertDirectorySchema(directories ?? []))
    closeReset()
  }

  const handleSaveInit = () => {
    openSave()
  }

  const handleSave = () => {
    save.mutate(roots, {
      onSuccess: () => notifications.show({ variant: "success", message: "Directories synced" }),
      onSettled: () => closeSave()
    })
  }

  return (
    <>
      <div className="flex items-center justify-end gap-2">
        <Button onClick={handleDirectoryCreate} variant="outline" leftSection={<FaPlus />} className="mr-8">
          Directory
        </Button>
        <Button onClick={handleResetInit} color="red">
          Reset
        </Button>
        <Button onClick={handleSaveInit}>
          Save
        </Button>
      </div>
      <Confirm
        opened={openedReset}
        onClose={closeReset}
        modalTitle="Reset"
        title="Reset directory structure"
        description="Are you sure you want to discard all changes?"
        onConfirm={handleReset}
      />
      <Confirm
        opened={openedSave}
        onClose={closeSave}
        modalTitle="Save"
        title="Save directory structure"
        description="Are you sure you want to save?"
        onConfirm={handleSave}
      />
    </>
  )
}

