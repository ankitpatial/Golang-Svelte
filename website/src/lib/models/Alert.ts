
export default interface Alert {
  id?: number;
  type: "info" | "success" | "warning" | "error";
  title?: string;
  body: string;
  confirmable?: boolean;
  onConfirm?: (data: any) => void;
}
