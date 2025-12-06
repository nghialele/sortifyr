import { MantineTheme } from "@mantine/core";
import { ReactNode } from "react";

// https://github.com/icflorescu/mantine-datatable/issues/651
declare module "mantine-datatable" {
  // https://icflorescu.github.io/mantine-datatable/type-definitions/

  interface Column<T> {
    accessor: keyof T | string;
    title?: string;
    width?: number;
    render?: (row: T) => ReactNode;
    sortable?: boolean;
    textAlign?: "left" | "right";
  }

  export interface DataTableProps<T> {
    idAccessor?: keyof T
    columns: Column<T>[]
    records: T[]

    minHeight?: number
    borderRadius?: "xs" | "sm" | "md" | "lg" | "xl"
    withTableBorder: boolean
    highlightOnHover?: boolean

    fetching?: boolean;
    customLoader?: ReactNode

    styles: {
      root: (theme: MantineTheme) => Record<string, string>,
      header: (theme: MantineTheme) => Record<string, string>,
    }
  }

  export function DataTable<T>(props: DataTableProps<T>): React.ReactNode;
}
