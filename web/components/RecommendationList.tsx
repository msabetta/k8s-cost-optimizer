import React from "react";

interface Recommendation {
  action: string;
  estimated_savings: number;
}

interface RecommendationListProps {
  recommendations: Recommendation[];
}

const RecommendationList: React.FC<RecommendationListProps> = ({ recommendations }) => {
  return (
    <ul className="divide-y divide-gray-200">
      {recommendations.map((rec, idx) => (
        <li key={idx} className="p-4 flex justify-between">
          <span>{rec.action}</span>
          <span className="font-semibold">{rec.estimated_savings.toFixed(2)} €</span>
        </li>
      ))}
    </ul>
  );
};

export default RecommendationList;
