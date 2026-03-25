import React from "react";

interface CostCardProps {
  title: string;
  value: number;
  unit?: string;
}

const CostCard: React.FC<CostCardProps> = ({ title, value, unit }) => {
  return (
    <div className="bg-white shadow rounded p-4 w-64">
      <h3 className="text-gray-500">{title}</h3>
      <p className="text-2xl font-bold">{value.toFixed(2)} {unit || "€"}</p>
    </div>
  );
};

export default CostCard;
