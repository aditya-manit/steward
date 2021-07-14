use crate::prelude::*;
use abscissa_core::{config, Command, FrameworkError, Options, Runnable};
use mongodb::{
    bson::{bson, doc},
    options::FindOptions,
    options::{ClientOptions, ServerAddress, StreamAddress},
    Client,
};

use crate::time_range::{MongoData, MongoTickWeights};

use futures::TryStreamExt;

#[derive(Command, Debug, Options)]
pub struct PredictionsCmd {
    /// To whom are we saying hello?
    #[options(free)]
    recipient: Vec<String>,
}

impl Runnable for PredictionsCmd {
    fn run(&self) {
        let config = APP.config();

        abscissa_tokio::run(&APP, async {
            // Parse a connection string into an options struct.
            let options = ClientOptions::builder()
            .hosts(vec![
                ServerAddress::Tcp {
                    host: "10.32.0.224".into(),
                    port: Some(27017),
                }
            ])
            .build();


            // Get a handle to the deployment.
            let client = Client::with_options(options).unwrap();

            // List the names of the databases in that deployment.
            for db_name in client.list_database_names(None, None).await.unwrap() {
                println!("{}", db_name);
            }
            let db = client.database("predictions");
            info!("Got database");

            for coll_name in db.list_collection_names(None).await.unwrap() {
                println!("collection: {}", coll_name);
            }

            // Get a handle to a collection in the database.
            let collection = db.collection::<MongoData>(
                "tick_range_predictions
        ",
            );

            let find_options = FindOptions::builder()
                .sort(doc! { "created_timestamp": -1 })
                .build();

            let mut sorted_predictions = collection.find(None, find_options).await.unwrap();

            if let Some(latest_prediction) = sorted_predictions.try_next().await.unwrap() {
                info!("lastet prediction:{:?}", latest_prediction)
            }
        })
        .unwrap_or_else(|e| {
            status_err!("executor exited with error: {}", e);
            std::process::exit(1);
        });
    }
}
