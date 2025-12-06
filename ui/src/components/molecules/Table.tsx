import { DataTable, DataTableProps } from "mantine-datatable";
import { LoadingSpinner } from "./LoadingSpinner"
import { useAutoAnimate } from '@formkit/auto-animate/react'

type Props<T> = Omit<
  DataTableProps<T>,
  "withTableBorder" | "styles"
> & Record<string, unknown>;

export const Table = <T,>(props: Props<T>) => {
  const [bodyRef] = useAutoAnimate<HTMLTableSectionElement>();

  return (
    <DataTable
      customLoader={<LoadingSpinner />}
      minHeight={180}
      backgroundColor="background.0"
      withTableBorder={false}
      textSelectionDisabled={true}
      styles={{
        root: (theme) => ({
          borderRadius: theme.radius.md,
        }),
        header: (theme) => ({
          background: theme.colors.secondary[1],
        }),
      }}
      bodyRef={bodyRef}
      {...props}
    />
  );
};
