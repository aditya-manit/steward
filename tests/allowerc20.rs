#![forbid(unsafe_code)]
#![warn(
    missing_docs,
    rust_2018_idioms,
    trivial_casts,
    unused_lifetimes,
    unused_qualifications
)]

use abscissa_core::testing::prelude::*;
use cellar_rebalancer::config::CellarRebalancerConfig;
use once_cell::sync::Lazy;

pub static RUNNER: Lazy<CmdRunner> = Lazy::new(|| CmdRunner::default());


#[test]
fn allow_erc20() {
    let mut runner = RUNNER.clone();
    runner
        .args(&[
            "tests/meconfig.toml",
            "allow-erc20",
            "--cellar-address=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
            "--address=0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
            "--amount",
            "50"
        ])
        .capture_stdout()
        .status()
        .expect_success();
}