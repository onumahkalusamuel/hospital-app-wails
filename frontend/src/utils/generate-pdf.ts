import { jsPDF } from "jspdf";

export function GeneratePDF(element: HTMLElement, filename: string, format: string = 'a4') {
  let pdf = new jsPDF({
    orientation: "p",
    unit: "mm",
    format,
  });

  // let pdf = new jsPDF('p', 'pt', 'a3');
  pdf.html(element, {
    margin: 1,
    image: { type: "jpeg", quality: 0.8 },
    html2canvas: { scale: .23, useCORS: true },
    callback: function (doc) {
      doc.save(filename);
    },
  });
}
