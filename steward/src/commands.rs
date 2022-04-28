//! Steward Subcommands
//!
//! This is where you specify the subcommands of your application.
//!
//! The default application comes with two subcommands:
//!
//! - `start`: launches the application
//! - `version`: print application version
//!
//! See the `impl Configurable` below for how to specify the path to the
//! application's configuration file.

mod config_cmd;
mod cosmos_to_eth;
mod deploy;
mod eth_to_cosmos;
mod keys;
mod orchestrator;
mod sign_delegate_keys;
mod start;

use self::{config_cmd::ConfigCmd, keys::KeysCmd, start::StartCmd};

use crate::config::StewardConfig;
use abscissa_core::{clap::Parser, Command, Configurable, FrameworkError, Runnable};
use std::{env, path::PathBuf};

/// Steward Configuration Filename
pub const CONFIG_FILE: &str = "steward.toml";

/// Steward Subcommands
#[derive(Command, Debug, Parser, Runnable)]
pub enum StewardCmd {
    #[clap(subcommand)]
    Keys(KeysCmd),
    /// Print default configurations
    PrintConfig(ConfigCmd),
    CosmosToEth(cosmos_to_eth::CosmosToEthCmd),
    #[clap(subcommand)]
    Deploy(deploy::DeployCmd),
    EthToCosmos(eth_to_cosmos::EthToCosmosCmd),
    #[clap(subcommand)]
    Orchestrator(orchestrator::OrchestratorCmd),
    SignDelegateKeys(sign_delegate_keys::SignDelegateKeysCmd),
    Start(StartCmd),
}

/// Entry point for the application. It needs to be a struct to allow using subcommands!
#[derive(Command, Debug, Parser)]
#[clap(author, about, version)]
pub struct EntryPoint {
    #[clap(subcommand)]
    cmd: StewardCmd,
    /// Enable verbose logging
    #[clap(short, long)]
    pub verbose: bool,
    /// Use the specified config file
    #[clap(short, long)]
    pub config: Option<String>,
}

impl Runnable for EntryPoint {
    fn run(&self) {
        self.cmd.run()
    }
}

/// This trait allows you to define how application configuration is loaded.
impl Configurable<StewardConfig> for EntryPoint {
    /// Location of the configuration file
    fn config_path(&self) -> Option<PathBuf> {
        // Check if the config file exists, and if it does not, ignore it.
        // If you'd like for a missing configuration file to be a hard error
        // instead, always return `Some(CONFIG_FILE)` here.
        let filename = self
            .config
            .as_ref()
            .map(PathBuf::from)
            .unwrap_or_else(|| CONFIG_FILE.into());

        let config_env_variable = option_env!("STW_ENV_VAR");

        if config_env_variable != None {
            let new_filename = self.config.as_ref().map(PathBuf::from).unwrap_or_else(|| {
                config_env_variable
                    .expect("Can't process config_env_variable")
                    .into()
            });
            println!("{}", new_filename.display());
            Some(new_filename)
        } else if filename.exists() {
            Some(filename)
        } else {
            None
        }
    }
    /// Apply changes to the config after it's been loaded, e.g. overriding
    /// values in a config file using command-line options.
    ///
    /// This can be safely deleted if you don't want to override config
    /// settings from command-line options.
    fn process_config(&self, config: StewardConfig) -> Result<StewardConfig, FrameworkError> {
        Ok(config)
    }
}
