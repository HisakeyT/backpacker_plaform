import { useEffect, useState } from "react";
import { Box, Typography } from "@mui/material";
import { getPublicTravelPlans } from "./repository";
import type { TravelPlan } from "./types";
import { PublicTravelPlanCard } from "./PublicTravelPlanCard";

type Props = {
  travelId: number;
};

const sortPlans = (plans: TravelPlan[]): TravelPlan[] => {
  return [...plans].sort((a, b) => {
    if (a.date !== b.date) return a.date < b.date ? -1 : 1;
    return a.sortOrder - b.sortOrder || a.id - b.id;
  });
};

const groupByDate = (plans: TravelPlan[]): [string, TravelPlan[]][] => {
  const map = new Map<string, TravelPlan[]>();
  for (const plan of plans) {
    const list = map.get(plan.date) ?? [];
    list.push(plan);
    map.set(plan.date, list);
  }
  return [...map.entries()];
};

export const PublicTravelPlanList = ({ travelId }: Props) => {
  const [plans, setPlans] = useState<TravelPlan[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPlans = async () => {
      try {
        const data = await getPublicTravelPlans(travelId);
        setPlans(data);
      } catch {
        setError("予定の取得に失敗しました");
      } finally {
        setIsLoading(false);
      }
    };

    fetchPlans();
  }, [travelId]);

  if (isLoading) return <Typography>Loading...</Typography>;
  if (error) return <Typography color="error">{error}</Typography>;
  if (plans.length === 0) {
    return (
      <Typography color="text.secondary">まだ予定がありません</Typography>
    );
  }

  const grouped = groupByDate(sortPlans(plans));

  return (
    <Box className="PublicTravelPlanList" sx={{ display: "grid", gap: 4 }}>
      {grouped.map(([date, dayPlans]) => (
        <Box key={date}>
          <Typography variant="h6" gutterBottom>
            {date}
          </Typography>
          <Box
            sx={{
              display: "grid",
              gap: 2,
              justifyContent: "center",
              gridTemplateColumns:
                "repeat(auto-fill, minmax(min(100%, 220px), 1fr))",
            }}
          >
            {dayPlans.map((plan) => (
              <PublicTravelPlanCard key={plan.id} plan={plan} />
            ))}
          </Box>
        </Box>
      ))}
    </Box>
  );
};
