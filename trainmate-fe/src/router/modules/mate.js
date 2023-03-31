import Layout from '@/layout'

const mateRouter = {
  path: '/mate',
  component: Layout,
  redirect: '/mate/jobs',
  name: 'Mate',
  meta: {
    title: 'Mate',
    icon: 'table'
  },
  children: [
    {
      path: 'task',
      component: () => import('@/views/mate/task'),
      name: 'Task',
      meta: { title: 'Task' }
    },
    {
      path: 'dataset',
      component: () => import('@/views/mate/dataset'),
      name: 'Dataset',
      meta: { title: 'Dataset' }
    },
    {
      path: 'Exp',
      component: () => import('@/views/mate/exp'),
      name: 'Experiments',
      meta: { title: 'Exp' }
    },
    {
      path: 'jobs',
      component: () => import('@/views/mate/jobs'),
      name: 'Jobs',
      meta: { title: 'Jobs' }
    }
    // {
    //   path: 'redirect',
    //   component: () => import('@/views/redirect/index'),
    //   name: 'Redirect',
    //   meta: { title: 'Redirect' }
    // }
  ]
}
export default mateRouter
