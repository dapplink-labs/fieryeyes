import express from 'express';
import { getCollections, getCollection } from "./services/collection.service";
import { getNft, getNfts } from "./services/nft.service";
import { getHolder, getHolders } from "./services/holder.service";
import { getChart } from "./services/chart.service";
import { getShadow } from "./services/radaer.service";
const realapikey = 'FgfW27uQdqXThfmVe-6y4Cq238e4X9-';
const router = express.Router();

// 健康检查接口
router.get('/health', (req, res) => {
  const { apikey } = req.headers;
  if (apikey !== realapikey) {
    return res.status(401).send('not authed');
  }
  return res.status(200).send('live');
});

router.post('/getCollections', async (req: any, res) => {
  const { apikey } = req.headers;
  if (apikey !== realapikey) {
    return res.status(401).send('not authed');
  }

  const { limit, skip, mint } = req.body;
  const result = await getCollections(skip, limit, mint);
  return res.status(200).send({
    data: result,
    success: !!result,
  });
});

router.get('/getCollections/:id', async (req: any, res) => {
  const { apikey } = req.headers;
  if (apikey !== realapikey) {
    return res.status(401).send('not authed');
  }

  const { id } = req.params;
  const result = await getCollection(id);
  return res.status(200).send({
    data: result,
    success: !!result,
  });
});

router.post('/nfts', async (req: any, res) => {
    const { apikey } = req.headers;
    if (apikey !== realapikey) {
      return res.status(401).send('not authed');
    }
  
    const { limit, skip } = req.body;
    const result = await getNfts(skip, limit);
    return res.status(200).send({
      data: result,
      success: !!result,
    });
  });
  
router.get('/nfts/:id', async (req: any, res) => {
    const { apikey } = req.headers;
    if (apikey !== realapikey) {
      return res.status(401).send('not authed');
    }
  
    const { id } = req.params;
    const result = await getNft(id);
    return res.status(200).send({
      data: result,
      success: !!result,
    });
  });

router.get('twitter/explore', async (req: any, res)=>{
    return res.status(200).send({
        data: ['tama', 'pink','dps'],
        success: true,
      });
})


router.post('/whaleHolders', async (req: any ,res) => {
    const { apikey } = req.headers;
    if (apikey !== realapikey) {
      return res.status(401).send('not authed');
    }
  
    const { limit, skip } = req.body;
    const result = await getHolders(skip, limit);
    return res.status(200).send({
      data: result,
      success: !!result,
    });
})

router.get('/whaleHolders/:id', async (req: any, res) => {
    const { apikey } = req.headers;
    if (apikey !== realapikey) {
      return res.status(401).send('not authed');
    }
  
    const { id } = req.params;
    const result = await getHolder(id);
    return res.status(200).send({
      data: result,
      success: !!result,
    });
  });

  router.post('/charts', async (req: any, res) => {
    const { apikey } = req.headers;
    if (apikey !== realapikey) {
      return res.status(401).send('not authed');
    }
  
    const { limit } = req.body;
    const result = await getChart(limit);
    return res.status(200).send({
      data: result,
      success: !!result,
    });
  });  

  router.post('/rodaer', async (req: any, res) => {
    const { apikey } = req.headers;
    if (apikey !== realapikey) {
      return res.status(401).send('not authed');
    }

    const result = await getShadow();
    return res.status(200).send({
      data: result,
      success: !!result,
    });
  });  

export default router;