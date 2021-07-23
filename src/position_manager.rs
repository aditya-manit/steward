use crate::{
    cellar_wrapper::CellarTickInfo,
    time_range::{TickWeight, TimeRange},
};

pub struct PositionManager {
    pub pair_id: ethers::types::U256,
    pub positions: Vec<Position>,
}

impl PositionManager {
    pub fn new() -> PositionManager {
     todo!()
    }
    pub fn compute_rebalance(&self, time_range: TimeRange) -> Vec<CellarTickInfo> {
        let mut cellars: Vec<CellarTickInfo> = Vec::new();
        for pos in &self.positions {
            let mut cellar_tick = CellarTickInfo {
                token_id: self.pair_id,
                tick_upper: pos.upper_tick,
                tick_lower: pos.lower_tick,
                weight: 0,
            };

            for tick in &time_range.tick_weights {
                let allocation = pos.within(tick);
                cellar_tick.weight = cellar_tick.weight + allocation;
            }
            cellars.push(cellar_tick);
        }

        return cellars;
    }
}

pub struct Position {
    pub upper_tick: i32,
    pub lower_tick: i32,
    pub weight: u32,
}

impl Position {
    pub fn within(&self, tick: &TickWeight) -> u32 {
        if tick.upper_bound <= self.upper_tick && tick.lower_bound >= self.upper_tick {
            return 1u32 * tick.weight;
        } else {
            return 0u32;
        }
    }
}